package container

import (
	"errors"
	"fmt"
)

var container map[string]interface{}
var params map[string]interface{}

//
// RegisterService ...
//
func RegisterService(name string, service interface{}) error {
	_, ok := container[name]

	if !ok {
		container[name] = service

		return nil
	}

	return errors.New("Service already registered")
}

//
// UnregisterService ...
//
func UnregisterService(name string) error {
	delete(container, name)

	return nil
}

//
// GetService ...
//
func GetService(name string) (interface{}, error) {
	service, ok := container[name]

	if !ok {
		return nil, errors.New("Could not find service with name [" + name + "]")
	}

	return service, nil
}

//
// GetServiceEx ...
//
func GetServiceEx(name string) interface{} {
	service, err := GetService(name)

	if err != nil {
		fmt.Println(name, err)
	}

	return service
}

//
// SetParam ...
//
func SetParam(name string, value interface{}) {
	params[name] = value
}

//
// GetParam ...
//
func GetParam(name string) interface{} {
	p, ok := params[name]

	if !ok {
		return nil
	}

	return p
}

func init() {
	container = make(map[string]interface{})
	params = make(map[string]interface{})
}
