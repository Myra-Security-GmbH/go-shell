package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	arg "github.com/alexflint/go-arg"
	"github.com/chzyer/readline"
	"myracloud.com/myra-shell/api"
	"myracloud.com/myra-shell/command"
	config "myracloud.com/myra-shell/config"
	"myracloud.com/myra-shell/console"
	"myracloud.com/myra-shell/container"
	"myracloud.com/myra-shell/context"
	"myracloud.com/myra-shell/eventHandler"
	"myracloud.com/myra-shell/output"
	"myracloud.com/myra-shell/util"

	_ "myracloud.com/myra-shell/command/impl"
)

var args struct {
	Configfile string `arg:"-c,help:Configfile to use"`
	Init       bool   `arg:"--init,help:Creates a default configuration file"`

	historyFile string
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

	args.historyFile = homePath + "/.myra-shell-history"
}

func createTemplateConfiguration() {
	_, err := os.Stat(args.Configfile)

	if err == nil {
		fmt.Println("File already exists")

		os.Exit(1)
	}

	fmt.Println("Creating config file")
	cfg := &config.Config{
		Endpoint: "https://api.myracloud.com",
		Language: "de",
		Login: []config.User{
			config.User{
				APIKey: "apiKey",
				Secret: "secret",
				User:   "username",
			},
		},
	}

	err = config.SaveConfigFile(args.Configfile, cfg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Config file successfully created")
	fmt.Println(": " + args.Configfile)

	os.Exit(0)
}

func readAvailableLogins(cfg *config.Config, ctx context.Context) {
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
		createTemplateConfiguration()
	}

	cfg, err := config.ReadConfigFile(args.Configfile)

	if err != nil {
		fmt.Printf("Error reading config file [%s]", args.Configfile)
		fmt.Println(err)
		os.Exit(1)
	}

	container.RegisterService("config", cfg)

	l, err := readline.NewEx(&readline.Config{
		Prompt:                 "",
		HistoryFile:            args.historyFile,
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

	readAvailableLogins(cfg, ctx)
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
