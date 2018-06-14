package impl

import (
	"errors"
	"fmt"

	"github.com/Myra-Security-GmbH/myra-shell/command"
	"github.com/Myra-Security-GmbH/myra-shell/context"
	"github.com/Myra-Security-GmbH/myra-shell/util"
)

//
// ClearCacheCommand ...
//
type ClearCacheCommand struct {
	*command.BaseCommand
}

//
// Completer ..,
//
func (cmd *ClearCacheCommand) Completer(
	ctx context.Container,
	input string,
) []string {
	return []string{}
}

//
// Execute ...
//
func (cmd *ClearCacheCommand) Execute(
	ctx context.Container,
	buffer *string,
) (uint, error) {
	connector := ctx.GetAPIConnector()
	clearCtx := ctx.FindSelection(context.AreaDomain)

	if len(cmd.GetParams()) < 2 {
		return 2, errors.New("Not enough parameters")
	} else if len(cmd.GetParams()) > 3 {
		return 2, errors.New("Too much parameters")
	} else if clearCtx == nil {
		return 1, errors.New("Command is not available in this context")
	}

	recursive := cmd.IsFlagSet("recursive")
	subdomainPattern := cmd.GetParameter(0, "")
	resourcePattern := cmd.GetParameter(1, "")

	if subdomainPattern == "." {
		_, err := connector.CacheClear(
			clearCtx.Identifier(),
			"",
			resourcePattern,
			recursive,
		)

		if err != nil {
			return 1, err
		}
	} else {
		filter := subdomainPattern
		types := &[]string{"A", "AAAA", "CNAME"}

		list, err := connector.DNSRecordList(
			clearCtx.Identifier(),
			&filter,
			types,
			true,
		)

		if err != nil {
			return 1, err
		}

		for _, ll := range list {
			connector.CacheClear(
				clearCtx.Identifier(),
				ll.Name,
				resourcePattern,
				recursive,
			)

			if cmd.IsFlagSet("verbose") {
				fmt.Println(ll.Name)
			}
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
			Description: util.NewTextTemplate("ss"),
			Name:        "verbose",
			ShortName:   "v",
		},
	}

	ClearCacheCommand := &ClearCacheCommand{
		BaseCommand: command.NewBaseCommand(
			command.CommandClearCache,
			argumentDefinitions,
			flagDefintions,
		),
	}

	command.RegisterCommand(ClearCacheCommand)
}

// func commandClearCache(c *Console, cmd *command.Command) uint {
// 	var identifier string
//
// 	usage := func() {
// 		fmt.Printf(
// 			"Usage %s [-r] <subdomainPattern> <resourcePattern>\n",
// 			command.CommandClearCache,
// 		)
// 	}
//
//
// 	clearCache := func(identifier string, name string, resourcePattern string, recursive bool) uint {
// 		cacheClearVOList, err := c.connector.CacheClear(
// 			identifier,
// 			name,
// 			resourcePattern,
// 			recursive,
// 		)
//
// 		if err != nil {
// 			output.Println(err)
// 			return 1
// 		}
//
// 		for _, vo := range cacheClearVOList {
// 			formatListing(
// 				interface{}(vo).(listing.Listable).RowListing(cmd),
// 				c.context.ListingDefintion(cmd),
// 				cmd.GetParameter(0, "*"),
// 			)
// 		}
//
// 		return 0
// 	}
//
// 	if subdomainPattern == "." {
// 		identifier = clearCtx.Identifier()
//
// 		ret := clearCache(
// 			identifier,
// 			"",
// 			resourcePattern,
// 			recursive,
// 		)
//
// 		if ret != 0 {
// 			return ret
// 		}
//
// 	} else {
// 		identifier = clearCtx.Name
//
// 		list, err := c.connector.DNSRecordList(
// 			ctx.Identifier(), nil, &[]string{"A", "AAAA", "CNAME"}, true,
// 		)
//
// 		if err != nil {
// 			output.Println(err)
// 			return 1
// 		}
//
// 		filteredList := make(map[string]bool)
//
// 		for _, record := range list {
// 			if fnmatch.Match(subdomainPattern, record.Name, fnmatch.FNM_IGNORECASE) {
// 				filteredList[record.Name] = true
// 			}
// 		}
//
// 		for recordName := range filteredList {
// 			ret := clearCache(
// 				identifier,
// 				recordName,
// 				resourcePattern,
// 				recursive,
// 			)
//
// 			if ret != 0 {
// 				return ret
// 			}
// 		}
// 	}
//
// 	return 0
// }
