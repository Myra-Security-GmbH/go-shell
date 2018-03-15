package storage

//
// Storage ...
//
type Storage interface {
	Get(name string) interface{}
	Set(name string, val interface{})
	Add(name string, item interface{}) error
	IsStructElemAvailable(name string, fieldName string, elem string) (interface{}, error)
	Len(name string) int
	Clear(name string)
	ClearAll()
}
