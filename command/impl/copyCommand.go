package impl

import (
	"errors"
	"fmt"

	"myracloud.com/myra-shell/api/vo"
	"myracloud.com/myra-shell/command"
	"myracloud.com/myra-shell/command/impl/listing"
	"myracloud.com/myra-shell/context"
	"myracloud.com/myra-shell/util"
)

//
// CopyCommand ...
//
type CopyCommand struct {
	*command.BaseCommand
}

//
// Completer ..,
//
func (cmd *CopyCommand) Completer(ctx context.Container, input string) []string {
	return []string{}
}

//
// Execute ...
//
func (cmd *CopyCommand) Execute(
	ctx context.Container,
	buffer *string,
) (uint, error) {
	tmp, err := command.FindCommand(command.CommandCd)

	if err != nil {
		return 1, err
	}

	cdCommand := tmp.(*ContextSwitchCommand)

	leftCtx := ctx.Clone()
	cdCommand.SetParams([]string{cmd.GetParameter(0, "")})
	cdCommand.Execute(leftCtx, buffer)

	rightCtx := ctx.Clone()
	cdCommand.SetParams([]string{cmd.GetParameter(1, "")})
	cdCommand.Execute(rightCtx, buffer)

	if leftCtx.GetSelection() != rightCtx.GetSelection() {
		return 1, errors.New("Copying data only in same context type")
	}

	tmp, err = command.FindCommand(command.CommandLs)

	if err != nil {
		return 1, err
	}

	var leftListing listing.RowListing

	lsCommand := tmp.(*listing.ListingCommand)
	leftListing, err = lsCommand.BuildListing(leftCtx, "*", "", false)

	if err != nil {
		return 1, err
	}

	rdors := rightCtx.FindSelection(context.AreaSubDomain, context.AreaDomain)

	for _, r := range leftListing {
		entity, ok := r.GetEntity().(vo.BaseEntityVOInterface)

		if ok {
			err = rightCtx.GetAPIConnector().SaveEntity(
				rdors.Identifier(),
				entity.ResetDatabaseState(),
			)

			if err != nil {
				return 1, err
			}

			if cmd.IsFlagSet("verbose") {
				fmt.Println(r.GetName())
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
			Name:        "source",
			Optional:    false,
		},
		&command.GenericArgumentDefinition{
			Description: util.NewTextTemplate(""),
			Example:     "4.5*",
			Name:        "target",
			Optional:    false,
		},
	}

	flagDefintions := []command.FlagDefintion{
		&command.GenericFlagDefinition{
			Description: util.NewTextTemplate("prompt before every removal"),
			Name:        "verbose",
			ShortName:   "v",
		},
	}

	cmd := &CopyCommand{
		BaseCommand: command.NewBaseCommand(
			command.CommandCopy,
			argumentDefinitions,
			flagDefintions,
		),
	}

	command.RegisterCommand(cmd)
}

// //
// // commandCopy ...
// //
// func commandCopy(c *Console, cmd *command.Command) uint {
// 	//ctx := c.context.FindSelection(context.AreaDomain, context.AreaSubDomain)
//
// 	if len(cmd.Params) != 2 {
// 		fmt.Println("Wrong amount of parameters")
// 		fmt.Println("Usage cp <from> <to>")
// 		return 1
// 	}
//
// 	pathData := strings.Split(cmd.Params[0], "/")
//
// 	filter := pathData[len(pathData)-1]
// 	pathData = pathData[:len(pathData)-1]
// 	ctxBackup := c.context
//
// 	if len(pathData) > 0 {
// 		c.SwitchTo(strings.Join(pathData, "/"))
// 	}
//
// 	selLeft := c.context.Selection
// 	identLeft := c.context.FindSelection(
// 		context.AreaSubDomain,
// 		context.AreaDomain,
// 	).Identifier()
//
// 	// copy data
//
// 	c.context = ctxBackup
// 	c.SwitchTo(cmd.Params[1])
// 	selRight := c.context.Selection
// 	identRight := c.context.FindSelection(
// 		context.AreaSubDomain,
// 		context.AreaDomain,
// 	).Identifier()
//
// 	if selLeft != selRight {
// 		fmt.Println(
// 			"Context to copy data to have to be the same as the context from",
// 		)
// 		c.context = ctxBackup
// 		return 1
// 	}
//
// 	var err error
// 	var list []interface{}
//
// 	switch selLeft {
// 	case context.AreaIPFilter:
// 		var tmp []vo.IPFilterVO
// 		tmp, err = c.connector.IPFilterList(identLeft, filter)
//
// 		for _, t := range tmp {
// 			list = append(list, t)
// 		}
// 		break
//
// 	case context.AreaRedirect:
// 		var tmp []vo.RedirectVO
// 		tmp, err = c.connector.RedirectList(identLeft, filter)
//
// 		for _, t := range tmp {
// 			list = append(list, t)
// 		}
// 		break
//
// 	case context.AreaCache:
// 		var tmp []vo.CacheSettingVO
// 		tmp, err = c.connector.CacheSettingList(identLeft, filter)
//
// 		for _, t := range tmp {
// 			list = append(list, t)
// 		}
// 		break
// 	}
//
// 	if err != nil {
// 		output.Println(err)
// 		return 1
// 	}
//
// 	_, verbose := cmd.Flags["v"]
//
// 	for _, itm := range list {
// 		bvo := itm.(vo.BaseEntityVOInterface).ResetDatabaseState()
//
// 		err := c.connector.SaveEntity(identRight, bvo)
//
// 		if err != nil {
// 			output.Println(err)
// 			return 1
// 		}
//
// 		if verbose {
// 			fmt.Printf("%+v\n", bvo)
// 		}
// 	}
//
// 	m := &sync.WaitGroup{}
// 	m.Add(1)
//
// 	observer.Publish(context.EventContextSwitch, context.SwitchContextEvent{
// 		Context:   ctxBackup,
// 		Direction: context.DirectionDown,
// 		WaitGroup: m,
// 	})
//
// 	m.Wait()
//
// 	return 0
// }
