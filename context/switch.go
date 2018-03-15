package context

import (
	"errors"

	"myracloud.com/myra-shell/storage"
)

//
// SwitchUp returns the parent context
//
func (c *myraContextContainer) SwitchUp() (Context, error) {
	if c.currentContext == nil || c.currentContext.parent == nil {
		return nil, errors.New("This is already on the root")
	}

	c.currentContext = c.currentContext.parent

	return c.currentContext, nil
}

//
// SwitchDown returns a new Context with the current Context
// as parent
//
func (c *myraContextContainer) SwitchDown(
	id uint64,
	name string,
	selection uint,
) (Context, error) {
	nc := &myraContext{
		Generic:   storage.NewGenericStorage(),
		id:        id,
		name:      name,
		selection: selection,
		parent:    c.currentContext,
	}

	err := c.GenericSubscriber.PublishEvent(
		EventContextSwitch,
		c,
		c.currentContext,
		nc,
	)

	if err != nil {
		return c.currentContext, err
	}

	c.currentContext = nc

	return nc, nil
}
