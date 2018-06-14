package context

import (
	"fmt"

	"github.com/Myra-Security-GmbH/myra-shell/storage"

	"github.com/logrusorgru/aurora"
)

//
// myraContext ...
//
type myraContext struct {
	*storage.Generic

	id        uint64
	name      string
	selection uint
	parent    *myraContext
}

//
// GetID ,,,
//
func (m *myraContext) GetID() uint64 {
	return m.id
}

//
// GetName ,,,
//
func (m *myraContext) GetName() string {
	return m.name
}

//
// GetSelection ,,,
//
func (m *myraContext) GetSelection() uint {
	return m.selection
}

//
// GetParent ,,,
//
func (m *myraContext) GetParent() Context {
	return m.parent
}

//
// BuildPrompt ,,,
//
func (m *myraContext) BuildPrompt(colorize bool) string {
	ret := "> "
	p := m

	for p != nil {
		sel := p.GetSelection()
		name := p.GetName()

		if !colorize && sel != AreaNone {
			ret = "/" + name + ret
		} else {
			if sel == AreaDomain {
				ret = aurora.Colorize("/", aurora.GreenFg|aurora.BoldFm).String() +
					aurora.Colorize(name, aurora.BlueFg|aurora.BoldFm).String() +
					ret
			} else if sel == AreaSubDomain {
				ret = aurora.Colorize("/", aurora.GreenFg|aurora.BoldFm).String() +
					aurora.Colorize(name, aurora.CyanFg|aurora.BoldFm).String() +
					ret
			} else if sel != AreaNone {
				ret = aurora.Colorize("/", aurora.GreenFg|aurora.BoldFm).String() +
					aurora.Colorize(name, aurora.BrownFg|aurora.BoldFm).String() +
					ret
			}
		}

		p = p.parent
	}

	return ret
}

//
// Identifier ...
//
func (m *myraContext) Identifier() string {
	if m.selection == AreaDomain {
		return "ALL:" + m.name
	}

	return m.name
}

//
// ToString ...
//
func (m *myraContext) String() string {
	return fmt.Sprintf(
		"[id=%d, selection=%d] %s",
		m.id, m.selection, m.name,
	)
}
