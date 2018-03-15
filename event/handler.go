package event

//
// Func ...
//
type Func func(source Subscriber, params ...interface{}) error

//
// Subscriber ...
//
type Subscriber interface {
	RegisterEvent(event string, handler ...Func) error
	UnregisterEvent(event string, handler ...Func) error
	PublishEvent(event string, params ...interface{}) error
}
