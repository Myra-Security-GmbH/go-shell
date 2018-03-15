package storage

import (
	"errors"
	"reflect"
)

//
// Generic ...
//
type Generic struct {
	data map[string]interface{}
}

//
// NewGenericStorage ...
//
func NewGenericStorage() *Generic {
	return &Generic{
		data: make(map[string]interface{}),
	}
}

//
// Get ...
//
func (s *Generic) Get(name string) interface{} {
	list, ok := s.data[name]

	if !ok {
		return nil
	}

	return list
}

//
// Set ...
//
func (s *Generic) Set(name string, val interface{}) {
	s.data[name] = val
}

//
// Add ...
//
func (s *Generic) Add(name string, item interface{}) error {
	l, ok := s.data[name].([]interface{})

	if ok {
		s.data[name] = append(l, item)

		return nil
	}

	return errors.New("Cannot add to non list value")
}

//
// IsStructElemAvailable ...
//
func (s *Generic) IsStructElemAvailable(
	name string,
	fieldName string,
	elem string,
) (interface{}, error) {
	list, ok := s.data[name]

	if !ok {
		return nil, errors.New("Cannot search in unset value")
	}

	kind := reflect.TypeOf(list).Kind()

	if kind == reflect.Ptr {
		kind = reflect.TypeOf(list).Elem().Kind()
	}

	if kind == reflect.Slice || kind == reflect.Array {
		len := reflect.ValueOf(list).Len()

		for i := 0; i < len; i++ {
			val := reflect.ValueOf(list).Index(i)

			if val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
				val = val.Elem()
			}

			if val.Kind() != reflect.Struct {
				return nil, errors.New("Element is not a of type struct")
			}

			if val.FieldByName(fieldName).String() == elem {
				return val.Interface(), nil
			}
		}
	}

	return nil, nil
}

//
// Len ...
//
func (s *Generic) Len(name string) int {
	list, ok := s.data[name].([]interface{})

	if !ok {
		return -1
	}

	return len(list)
}

//
// Clear ...
//
func (s *Generic) Clear(name string) {
	delete(s.data, name)
}

//
// ClearAll ...
//
func (s *Generic) ClearAll() {
	s.data = make(map[string]interface{})
}
