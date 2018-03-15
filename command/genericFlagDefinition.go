package command

import (
	"bytes"
	text "text/template"
)

//
// GenericFlagDefinition ...
//
type GenericFlagDefinition struct {
	ShortName    string
	Name         string
	Description  *text.Template
	Example      string
	DefaultValue interface{}
}

//
// GetShortName ...
//
func (fd *GenericFlagDefinition) GetShortName() string {
	return fd.ShortName
}

//
// GetName ...
//
func (fd *GenericFlagDefinition) GetName() string {
	return fd.Name
}

//
// GetDescription ...
//
func (fd *GenericFlagDefinition) GetDescription(selection uint) string {
	var bb = new(bytes.Buffer)

	err := fd.Description.Execute(bb, struct {
		Selection uint
	}{
		Selection: selection,
	})

	if err != nil {
		return err.Error()
	}

	return bb.String()
}

//
// GetExample ...
//
func (fd *GenericFlagDefinition) GetExample() string {
	return fd.Example
}

//
// GetDefaultValue ...
//
func (fd *GenericFlagDefinition) GetDefaultValue() interface{} {
	return fd.DefaultValue
}
