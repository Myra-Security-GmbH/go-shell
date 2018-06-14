package command

import (
	"github.com/Myra-Security-GmbH/myra-shell/context"
)

//
// Command ...
//
type Command interface {
	GetParameter(idx int, defaultValue string) string
	GetCommand() string
	GetParams() []string
	GetFlags() map[string]bool
	IsFlagSet(flag ...string) bool
	SetParams(params []string)
	SetFlags(flags map[string]bool)

	Execute(ctx context.Container, buffer *string) (uint, error)
	Completer(ctx context.Container, input string) []string
}
