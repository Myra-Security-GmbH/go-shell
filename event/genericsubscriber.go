package event

//
// GenericSubscriber ...
//
type GenericSubscriber struct {
	handler map[string][]Func
}

//
// NewGenericSubscriber ...
//
func NewGenericSubscriber() *GenericSubscriber {
	return &GenericSubscriber{
		handler: make(map[string][]Func),
	}
}

//
// PublishEvent ...
//
func (c *GenericSubscriber) PublishEvent(event string, params ...interface{}) error {
	var err error

	list, ok := c.handler[event]

	if !ok {
		return nil
	}

	for _, fn := range list {
		err = fn(c, params...)

		if err != nil {
			break
		}
	}

	return err
}

//
// RegisterEvent ...
//
func (c *GenericSubscriber) RegisterEvent(
	event string,
	handler ...Func,
) error {
	_, ok := c.handler[event]

	if !ok {
		c.handler[event] = []Func{}
	}

	for _, h := range handler {
		c.handler[event] = append(c.handler[event], h)
	}

	return nil
}

//
// UnregisterEvent ...
//
func (c *GenericSubscriber) UnregisterEvent(
	event string,
	handler ...Func,
) error {
	delete(c.handler, event)

	return nil
}
