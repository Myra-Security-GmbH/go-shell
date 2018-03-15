package command

//
// ParamDefinition ...
//
type ParamDefinition interface {
	GetName() string
	GetDescription(selection uint) string
	GetExample() string
}

//
// ArgumentDefintion ...
//
type ArgumentDefintion interface {
	ParamDefinition

	IsOptional() bool
}

//
// FlagDefintion ...
//
type FlagDefintion interface {
	ParamDefinition

	GetShortName() string
	GetDefaultValue() interface{}
}
