package context

import (
	"errors"
	"strings"

	"github.com/Myra-Security-GmbH/myra-shell/api"
	"github.com/Myra-Security-GmbH/myra-shell/event"
	"github.com/Myra-Security-GmbH/myra-shell/storage"
)

//
// NewMyraContainer ...
//
func NewMyraContainer(connector api.API) Container {
	return &myraContextContainer{
		GenericSubscriber: event.NewGenericSubscriber(),
		apiConnector:      connector,
		currentContext: &myraContext{
			Generic:   storage.NewGenericStorage(),
			id:        0,
			name:      "NONE",
			selection: AreaNone,
		},
	}
}

//
// myraContextContainer ...
//
type myraContextContainer struct {
	*event.GenericSubscriber

	currentContext *myraContext
	apiConnector   api.API
}

func (c *myraContextContainer) Reset() {
	p := c.currentContext

	for p != nil {
		if p.parent == nil {
			c.currentContext = p
			break
		}

		p = p.parent
	}
}

//
// GetAPIConnector ...
//
func (c *myraContextContainer) GetAPIConnector() api.API {
	return c.apiConnector
}

//
// Clone ...
//
func (c *myraContextContainer) Clone() Container {
	var nc *myraContext
	pnc := nc
	p := c.currentContext

	for p != nil {
		tmpc := &myraContext{
			Generic:   p.Generic,
			id:        p.GetID(),
			name:      p.GetName(),
			selection: p.GetSelection(),
		}

		if pnc == nil {
			pnc = tmpc
		} else {
			pnc.parent = tmpc
			pnc = pnc.parent
		}

		if nc == nil {
			nc = pnc
		}

		p = p.parent
	}

	ac := c.GetAPIConnector()

	nac, _ := api.NewMyraAPI(
		ac.GetAPIKey(),
		ac.GetSecret(),
		ac.GetEndpoint(),
		ac.GetLanguage(),
	)

	return &myraContextContainer{
		GenericSubscriber: event.CloneSubscriber(c.GenericSubscriber).(*event.GenericSubscriber),
		apiConnector:      nac,
		currentContext:    nc,
	}
}

//
// GetID ...
//
func (c *myraContextContainer) GetID() uint64 {
	if c.currentContext != nil {
		return c.currentContext.id
	}

	return 0
}

//
// GetName ...
//
func (c *myraContextContainer) GetName() string {
	if c.currentContext != nil {
		return c.currentContext.name
	}

	return ""
}

//
// GetSelection ...
//
func (c *myraContextContainer) GetSelection() uint {
	if c.currentContext != nil {
		return c.currentContext.selection
	}

	return 0
}

//
// GetParent ...
//
func (c *myraContextContainer) GetParent() Context {
	if c.currentContext != nil {
		return Context(c.currentContext.parent)
	}

	return nil
}

//
// ToString ...
//
func (c *myraContextContainer) String() string {
	var ret string

	p := c.currentContext

	cc := 0
	for p != nil {
		if cc > 0 {
			ret += strings.Repeat(" ", cc*2) + "└─>"
		}

		ret += p.String() + "\n"

		cc++

		p = p.parent
	}

	return ret
}

//
// FindSelection ...
//
func (c *myraContextContainer) FindSelection(selection ...uint) Context {
	p := c.currentContext

	for p != nil {
		for _, s := range selection {
			if p.selection == s {
				return p
			}
		}

		p = p.parent
	}

	return nil
}

//
// BuildPrompt ...
//
func (c *myraContextContainer) BuildPrompt(colorize bool) string {
	if c.currentContext != nil {
		return c.currentContext.BuildPrompt(colorize)
	}

	return "> "
}

func (c *myraContextContainer) Identifier() string {
	if c.currentContext != nil {
		return c.currentContext.Identifier()
	}

	return ""
}

//
// Get ...
//
func (c *myraContextContainer) Get(name string) interface{} {
	if c.currentContext != nil {
		return c.currentContext.Get(name)
	}

	return nil
}

//
// Set ...
//
func (c *myraContextContainer) Set(name string, val interface{}) {
	if c.currentContext != nil {
		c.currentContext.Set(name, val)
	}
}

//
// Add ...
//
func (c *myraContextContainer) Add(name string, item interface{}) error {
	if c.currentContext != nil {
		return c.currentContext.Add(name, item)
	}

	return errors.New("No context available for storage operation")
}

//
// IsStructElemAvailable ...
//
func (c *myraContextContainer) IsStructElemAvailable(
	name string,
	fieldName string,
	elem string,
) (interface{}, error) {
	if c.currentContext != nil {
		return c.currentContext.IsStructElemAvailable(name, fieldName, elem)
	}

	return nil, errors.New("No context available for storage operation")
}

//
// Len ...
//
func (c *myraContextContainer) Len(name string) int {
	if c.currentContext != nil {
		return c.currentContext.Len(name)
	}

	return 0
}

//
// Clear ...
//
func (c *myraContextContainer) Clear(name string) {
	if c.currentContext != nil {
		c.currentContext.Clear(name)
	}
}

//
// ClearAll ...
//
func (c *myraContextContainer) ClearAll() {
	if c.currentContext != nil {
		c.currentContext.ClearAll()
	}
}
