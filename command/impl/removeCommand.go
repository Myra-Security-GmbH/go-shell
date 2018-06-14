package impl

import (
	"errors"
	"fmt"

	"github.com/Myra-Security-GmbH/myra-shell/command"
	"github.com/Myra-Security-GmbH/myra-shell/command/impl/listing"
	"github.com/Myra-Security-GmbH/myra-shell/context"
	"github.com/Myra-Security-GmbH/myra-shell/util"
)

//
// RemoveCommand ...
//
type RemoveCommand struct {
	*command.BaseCommand
}

//
// Completer ..,
//
func (cmd *RemoveCommand) Completer(ctx context.Container, input string) []string {
	return []string{}
}

//
// Execute ...
//
func (cmd *RemoveCommand) Execute(ctx context.Container, buffer *string) (uint, error) {
	if cmd.GetParameter(0, "") == "" {
		return 1, errors.New("Usage rm <patter for source to remove>")
	}

	verbose := cmd.IsFlagSet("verbose")

	dors := ctx.FindSelection(context.AreaDomain, context.AreaSubDomain)

	if dors == nil {
		return 1, errors.New("Create command not available in this context")
	}

	connector := ctx.GetAPIConnector()

	lscmd, err := command.FindCommand(command.CommandLs)

	if err != nil {
		return 1, err
	}

	var filter = cmd.GetParameter(0, "*")
	var apiFilter = cmd.GetParameter(0, "")

	rows, err := lscmd.(*listing.ListingCommand).BuildListing(
		ctx, filter, apiFilter, false,
	)

	if err != nil {
		return 1, err
	}

	for _, r := range rows {
		err = connector.RemoveEntity(dors.Identifier(), r.GetEntity())

		if err != nil {
			return 1, err
		}

		if verbose {
			fmt.Println(r.GetID(), r.GetName())
		}
	}

	return 0, nil
}

func init() {
	argumentDefinitions := []command.ArgumentDefintion{
		&command.GenericArgumentDefinition{
			Description: util.NewTextTemplate(""),
			Example:     "4.5*",
			Name:        "filter",
			Optional:    true,
		},
	}

	flagDefintions := []command.FlagDefintion{
		&command.GenericFlagDefinition{
			Description: util.NewTextTemplate("prompt before every removal"),
			Name:        "verbose",
			ShortName:   "v",
		},
	}

	removeCommand := &RemoveCommand{
		BaseCommand: command.NewBaseCommand(
			command.CommandRm,
			argumentDefinitions,
			flagDefintions,
		),
	}

	command.RegisterCommand(removeCommand)
}
