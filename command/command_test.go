// +build !testing

package command

import (
	"errors"
	"fmt"
	"testing"

	"myracloud.com/myra-shell/context"

	"github.com/stretchr/testify/require"
)

type MockCommand struct {
	*BaseCommand
}

func (cmd *MockCommand) Execute(ctx context.Container, buffer *string) (uint, error) {
	*buffer = "MUH"
	return 0, nil
}

func (cmd *MockCommand) Completer(ctx context.Container, input string) []string {
	return []string{}
}

func init() {
	flagDef := []FlagDefintion{
		&GenericFlagDefinition{
			Name:      "test",
			ShortName: "t",
		},
		&GenericFlagDefinition{
			Name:      "all",
			ShortName: "a",
		},
		&GenericFlagDefinition{
			Name:      "interactive",
			ShortName: "i",
		},
		&GenericFlagDefinition{
			Name:      "verbose",
			ShortName: "v",
		},
		&GenericFlagDefinition{
			Name:      "long",
			ShortName: "l",
		},
	}

	cmd := &MockCommand{
		BaseCommand: NewBaseCommand(CommandLs, nil, flagDef),
	}
	RegisterCommand(cmd)
}

type editorTestCase struct {
	err     error
	input   string
	command string
	params  []string
	flags   []string
}

func TestParseEditorCommand(t *testing.T) {
	cases := []editorTestCase{
		editorTestCase{
			input:   "ls test 123 -t",
			command: "ls",
			params:  []string{"test", "123"},
			flags:   []string{"t"},
		},

		editorTestCase{
			input:   "ls -ta --interactive tttt xxx",
			command: "ls",
			params:  []string{"tttt", "xxx"},
			flags:   []string{"t", "a", "interactive"},
		},

		editorTestCase{
			input:   `ls -v -l 1111 2222`,
			command: "ls",
			params:  []string{"1111", "2222"},
			flags:   []string{"v", "l"},
		},

		editorTestCase{
			input:   `ls -v -l "1111 2222"`,
			command: "ls",
			params:  []string{"1111 2222"},
			flags:   []string{"v", "l"},
		},

		editorTestCase{
			input: `ls -v -l "1111 2222`,
			err:   errors.New("Missing closing doublequote"),
		},
	}

	for _, c := range cases {
		cmd, err := ParseEditorCommand(c.input)

		require.Equal(t, c.err, err, fmt.Sprintf("%+v", c))

		if err == nil {
			require.Equal(t, c.command, cmd.GetCommand())
			require.Equal(t, c.params, cmd.GetParams())

			for _, f := range c.flags {
				require.True(
					t,
					cmd.IsFlagSet(f),
					fmt.Sprintf("Flag [%s] does not exist on:\n %+v", f, c),
				)
			}

			buffer := ""
			cmd.Execute(nil, &buffer)
			require.Equal(t, "MUH", buffer, fmt.Sprintf("%+v", c))
		}
	}
}
