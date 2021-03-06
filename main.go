package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Myra-Security-GmbH/myra-shell/api"
	"github.com/Myra-Security-GmbH/myra-shell/command"
	_ "github.com/Myra-Security-GmbH/myra-shell/command/impl"
	"github.com/Myra-Security-GmbH/myra-shell/console"
	"github.com/Myra-Security-GmbH/myra-shell/container"
	"github.com/Myra-Security-GmbH/myra-shell/context"
	"github.com/Myra-Security-GmbH/myra-shell/eventHandler"
	"github.com/Myra-Security-GmbH/myra-shell/output"
	"github.com/Myra-Security-GmbH/myra-shell/util"

	config "github.com/Myra-Security-GmbH/myra-shell/config"
	arg "github.com/alexflint/go-arg"
	"github.com/chzyer/readline"
)

var args struct {
	Configfile string `arg:"-c" help:"Configfile to use"`
	Init       bool   `arg:"--init" help:"Creates a default configuration file"`

	HistoryFile string
}

func defineDefaults() {
	homePath, homeExist := os.LookupEnv("HOME")

	if !homeExist {
		fmt.Println("History disabled. HOME environment variable is not set.")
	}

	args.Configfile = util.GetFirstPathMatches([]string{
		"./config.yml",
		homePath + "/config.yml",
	})

	if len(args.Configfile) == 0 {
		args.Configfile = "./config.yml"
	}

	args.HistoryFile = homePath + "/.myra-shell-history"
}

func createTemplateConfiguration() error {
	_, err := os.Stat(args.Configfile)

	if err == nil {
		return nil
	}

	fmt.Println("Creating config file")

	username, apiKey, secret := promptCredentials()

	cfg := &config.Config{
		Endpoint: "https://api.myracloud.com",
		Language: "de",
		Login: []config.User{
			config.User{
				APIKey: apiKey,
				Secret: secret,
				User:   username,
			},
		},
	}

	err = config.SaveConfigFile(args.Configfile, cfg)

	if err != nil {
		return err
	}

	fmt.Printf("Config file %s successfully created\n", args.Configfile)
	return nil
}

func promptCredentials() (string, string, string) {
	fmt.Print("Username: ")
	username := bufio.NewScanner(os.Stdin)
	username.Scan()

	fmt.Print("API-Key: ")
	apiKey := bufio.NewScanner(os.Stdin)
	apiKey.Scan()

	fmt.Print("Secret: ")
	secret := bufio.NewScanner(os.Stdin)
	secret.Scan()

	return username.Text(), apiKey.Text(), secret.Text()
}

func readAvailableLogins(ctx context.Context, cfg *config.Config) {
	// stat, _ := os.Stdin.Stat()
	// isNonInterface := (stat.Mode()&os.ModeCharDevice == 0)

	var loginList []console.Login

	for _, k := range cfg.Login {
		loginList = append(loginList, console.Login{
			APIKey: k.APIKey,
			Secret: k.Secret,
			Name:   k.User,
		})
	}

	ctx.Set("userList", loginList)

	// if !isNonInterface && store.Len(storage.StorageTypeUserList) == 1 {
	// 	user := store.GetList(storage.StorageTypeUserList)[0].(console.Login)
	// 	consoleHandler.SwitchContextDown(user.Name, context.AreaUser)
	// }
}

//
// main
//
func main() {
	defineDefaults()
	arg.MustParse(&args)

	if args.Init {
		err := createTemplateConfiguration()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cfg, err := config.ReadConfigFile(args.Configfile)
	if err != nil {
		if os.IsNotExist(err) {
			err = createTemplateConfiguration()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			cfg, err = config.ReadConfigFile(args.Configfile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("Error reading config file [%s]", args.Configfile)
			fmt.Println(err)
			os.Exit(1)
		}
	}

	container.RegisterService("config", cfg)

	l, err := readline.NewEx(&readline.Config{
		Prompt:                 "",
		HistoryFile:            args.HistoryFile,
		InterruptPrompt:        "^C",
		EOFPrompt:              "exit",
		DisableAutoSaveHistory: true,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.SetOutput(l.Stderr())

	connector, err := api.NewMyraAPI(
		"", "", cfg.Endpoint, cfg.Language,
	)

	ctx := context.NewMyraContainer(connector)
	ctx.RegisterEvent(context.EventContextSwitch, eventHandler.SwitchContextEvent)

	if err != nil {
		panic(err)
	}

	container.RegisterService("context", ctx)
	container.RegisterService("api", connector)

	readAvailableLogins(ctx, cfg)
	l.Config.AutoComplete = command.BuildCompleter()

	var returnValue uint
	var buffer string

	for {
		l.SetPrompt(ctx.BuildPrompt(true))

		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		commands := strings.Split(line, "&&")

		for _, cmdLine := range commands {
			cmdLine = strings.TrimSpace(cmdLine)
			cmd, err := command.ParseEditorCommand(cmdLine)

			if err != nil {
				output.Println(err)

				fmt.Println("ERR")
				returnValue = 1
			} else if cmd != nil {
				buffer = ""
				returnValue, err = cmd.Execute(ctx, &buffer)

				if err != nil {
					fmt.Println(err)
				}

				fmt.Print(buffer)
			}

			if returnValue == 0 {
				if l.Config.HistoryFile != "" {
					l.SaveHistory(cmdLine)
				}
			} else {
				break
			}
		}
	}
}
