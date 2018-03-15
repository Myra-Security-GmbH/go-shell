package command

import (
	"bytes"
	text "text/template"
)

//
// GenericArgumentDefinition ...
//
type GenericArgumentDefinition struct {
	Name        string
	Description *text.Template
	Example     string
	Optional    bool
}

//
// GetName ...
//
func (fd *GenericArgumentDefinition) GetName() string {
	return fd.Name
}

//
// GetDescription ...
//
func (fd *GenericArgumentDefinition) GetDescription(selection uint) string {
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
func (fd *GenericArgumentDefinition) GetExample() string {
	return fd.Example
}

//
// IsOptional ...
//
func (fd *GenericArgumentDefinition) IsOptional() bool {
	return fd.Optional
}
