package impl

import (
	"errors"
	"fmt"
	"strings"

	"myracloud.com/myra-shell/api/vo"
	"myracloud.com/myra-shell/command"
	"myracloud.com/myra-shell/console"
	"myracloud.com/myra-shell/context"
)

//
// ContextSwitchCommand ...
//
type ContextSwitchCommand struct {
	*command.BaseCommand
}

//
// Completer ..,
//
func (cmd *ContextSwitchCommand) Completer(
	ctx context.Container,
	input string,
) []string {
	var ret []string

	switch ctx.GetSelection() {
	case context.AreaNone:
		ll, ok := ctx.Get("userList").([]console.Login)

		if !ok {
			break
		}

		for _, l := range ll {
			ret = append(ret, l.Name)
		}

	case context.AreaUser:
		ll, ok := ctx.Get("domainList").([]vo.DomainVO)

		if !ok {
			break
		}

		for _, l := range ll {
			ret = append(ret, l.Name)
		}

	case context.AreaDomain:
		ll, ok := ctx.Get("subDomainList").([]vo.DNSRecordVO)

		if !ok {
			break
		}

		ret = append(ret,
			context.ContextCacheSettings,
			context.ContextIPFilter,
			context.ContextRedirect,
			context.ContextSettings,
			context.ContextStatistics,
			context.ContextErrorPages,
			context.ContextSsl,
		)

		for _, l := range ll {
			ret = append(ret, l.Name)
		}

	case context.AreaSubDomain:
		ret = append(ret,
			context.ContextCacheSettings,
			context.ContextIPFilter,
			context.ContextRedirect,
			context.ContextSettings,
			context.ContextStatistics,
			context.ContextSsl,
		)

	default:
		fmt.Println(ctx)
	}

	return ret
}

//
// Execute ...
//
func (cmd *ContextSwitchCommand) Execute(ctx context.Container, buffer *string) (uint, error) {
	if len(cmd.GetParams()) < 1 {
		return 1, errors.New("Missing argument")
	}

	var path = cmd.GetParameter(0, "")
	var contextList []string

	if path[0:1] == "/" {
		contextList = strings.Split(path[1:], "/")
		ctx.Reset()
	} else {
		contextList = strings.Split(path, "/")
	}

	var targetArea uint
	var targetID uint64

	for _, targetContextName := range contextList {
		targetArea = 0
		targetID = 0

		if targetContextName == context.ContextUp {
			_, err := ctx.SwitchUp()

			if err != nil {
				return 1, err
			}
			continue
		}

		if targetContextName == context.ContextSelf {
			continue
		}

		sel := ctx.GetSelection()

		switch {
		case (sel == context.AreaNone):
			u, err := ctx.IsStructElemAvailable("userList", "Name", targetContextName)

			if err != nil {
				return 1, err
			}

			if u != nil {
				targetArea = context.AreaUser
			}
			break

		case (sel == context.AreaUser):
			domain, err := ctx.IsStructElemAvailable("domainList", "Name", targetContextName)

			if err != nil {
				return 1, err
			}

			if domain != nil {
				targetArea = context.AreaDomain
				targetID = *domain.(vo.DomainVO).ID
			}
			break

		case (sel == context.AreaDomain || sel == context.AreaSubDomain):
			if targetContextName == context.ContextIPFilter {
				targetArea = context.AreaIPFilter
			} else if targetContextName == context.ContextCacheSettings {
				targetArea = context.AreaCache
			} else if targetContextName == context.ContextRedirect {
				targetArea = context.AreaRedirect
			} else if targetContextName == context.ContextSettings {
				targetArea = context.AreaSettings
			} else if targetContextName == context.ContextStatistics {
				targetArea = context.AreaStatistics
			} else if targetContextName == context.ContextErrorPages {
				targetArea = context.AreaErrorPage
			} else if targetContextName == context.ContextSsl {
				targetArea = context.AreaSsl
			} else if sel == context.AreaDomain {
				record, err := ctx.IsStructElemAvailable(
					"subDomainList", "Name", targetContextName,
				)

				if err != nil {
					return 1, err
				}

				if record != nil {
					targetArea = context.AreaSubDomain
					targetID = *record.(vo.DNSRecordVO).ID
				}
			}
			break
		}

		if targetArea == 0 {
			return 1, errors.New("Cannot switch context to: " + path)
		}
		_, err := ctx.SwitchDown(targetID, targetContextName, targetArea)

		if err != nil {
			return 1, err
		}
	}

	return 0, nil
}

func init() {
	cmd := &ContextSwitchCommand{
		BaseCommand: command.NewBaseCommand(command.CommandCd, nil, nil),
	}

	command.RegisterCommand(cmd)
}
