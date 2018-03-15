package output

//
// Help ...
//
type Help interface {
	Help() string
	Command() string
	Params() []HelpParam
	GetParamLen() (int, int)
}

//
// HelpParam ...
//
type HelpParam struct {
	Name         string
	LongName     string
	Description  string
	DefaultValue string
	Example      string
	Required     bool
	Flag         bool
}

type helpMessage struct {
	command     string
	description string
	params      []HelpParam
}

//
// Help returns the message.
//
func (e *helpMessage) Help() string {
	return e.description
}

//
// Command returns the command.
//
func (e *helpMessage) Command() string {
	return e.command
}

//
// Params returns a list of options and arguments.
//
func (e *helpMessage) Params() []HelpParam {
	return e.params
}

//
// GetParamLen returns length of arguments and options.
//
func (e *helpMessage) GetParamLen() (argLen int, opLen int) {
	for _, p := range e.Params() {
		if p.Flag {
			opLen++
		} else {
			argLen++
		}
	}

	return
}

//
// NewHelp ...
//
func NewHelp(command string, help string, params ...HelpParam) Help {
	return &helpMessage{
		command:     command,
		description: help,
		params:      params,
	}
}
