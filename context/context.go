package context

import (
	"fmt"

	"myracloud.com/myra-shell/api"
	"myracloud.com/myra-shell/event"
	"myracloud.com/myra-shell/storage"
)

//
// Context ...
//
type Context interface {
	fmt.Stringer
	storage.Storage

	GetID() uint64
	GetName() string
	GetSelection() uint
	GetParent() Context
	BuildPrompt(colorize bool) string
	Identifier() string
}

//
// Container ...
//
type Container interface {
	Context
	event.Subscriber

	SwitchUp() (Context, error)
	SwitchDown(id uint64, name string, selection uint) (Context, error)
	FindSelection(selection ...uint) Context
	GetAPIConnector() api.API
	Reset()
	Clone() Container
}
