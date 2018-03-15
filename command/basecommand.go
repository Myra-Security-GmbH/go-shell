package command

//
// BaseCommand ...
//
type BaseCommand struct {
	command string
	params  []string
	flags   map[string]bool

	flagDefintions     []FlagDefintion
	argumentDefintions []ArgumentDefintion
}

//
// NewBaseCommand ...
//
func NewBaseCommand(cmd string, argDefs []ArgumentDefintion, flagDefs []FlagDefintion) *BaseCommand {
	return &BaseCommand{
		flags:              make(map[string]bool),
		command:            cmd,
		argumentDefintions: argDefs,
		flagDefintions:     flagDefs,
	}
}

//
// GetParameter ...
//
func (cmd *BaseCommand) GetParameter(idx int, defaultValue string) string {
	if len(cmd.params) > idx {
		return cmd.params[idx]
	}

	return defaultValue
}

//
// GetCommand ...
//
func (cmd *BaseCommand) GetCommand() string {
	return cmd.command
}

//
// GetParams ...
//
func (cmd *BaseCommand) GetParams() []string {
	return cmd.params
}

//
// GetFlags ...
//
func (cmd *BaseCommand) GetFlags() map[string]bool {
	return cmd.flags
}

//
// IsFlagSet ...
//
func (cmd *BaseCommand) IsFlagSet(flag ...string) bool {
	for _, def := range cmd.flagDefintions {
		for _, f := range flag {
			if def.GetName() == f || def.GetShortName() == f {
				ret, ok := cmd.flags[def.GetName()]

				if ok {
					return ret
				}
			}
		}
	}

	return false
}

//
// SetParams ...
//
func (cmd *BaseCommand) SetParams(params []string) {
	cmd.params = params
}

//
// SetFlags ...
//
func (cmd *BaseCommand) SetFlags(flags map[string]bool) {
	for _, def := range cmd.flagDefintions {
		cmd.flags[def.GetName()] = false

		for name := range flags {
			if def.GetName() == name || def.GetShortName() == name {
				cmd.flags[def.GetName()] = true
				break
			}
		}
	}
}
