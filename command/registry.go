package command

import (
	"errors"

	"myracloud.com/myra-shell/container"
	"myracloud.com/myra-shell/context"

	"github.com/chzyer/readline"
)

var commandRegistry = make(map[string]Command)

//
// FindCommand ...
//
func FindCommand(cmd string) (Command, error) {
	ret, ok := commandRegistry[cmd]

	if ok {
		return ret, nil
	}

	return nil, errors.New("Invalid command")
}

//
// RegisterCommand ...
//
func RegisterCommand(cmd Command) {
	commandRegistry[cmd.GetCommand()] = cmd
}

//
// BuildCompleter ...
//
func BuildCompleter() *readline.PrefixCompleter {
	ret := readline.NewPrefixCompleter()

	for str, cmd := range commandRegistry {
		itm := readline.PcItem(str)

		completer := func(cmd Command) func(input string) []string {
			return func(input string) []string {
				ctx := container.GetServiceEx("context").(context.Container)

				return cmd.Completer(ctx, input)
			}
		}

		itm.Children = append(itm.Children, readline.PcItemDynamic(
			completer(cmd),
		))

		ret.Children = append(ret.Children, itm)
	}

	return ret
}
