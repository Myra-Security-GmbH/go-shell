package impl

import (
	"myracloud.com/myra-shell/command"
	"myracloud.com/myra-shell/command/impl/listing"
	"myracloud.com/myra-shell/util"
)

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
			Description: util.NewTextTemplate("Longmode shows additional data."),
			Name:        "long",
			ShortName:   "l",
		},
	}

	listingCommand := &listing.ListingCommand{
		BaseCommand: command.NewBaseCommand(
			command.CommandLs,
			argumentDefinitions,
			flagDefintions,
		),
	}

	command.RegisterCommand(listingCommand)
}
