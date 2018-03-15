package tag

import (
	"reflect"
	"strconv"
	"strings"

	"myracloud.com/myra-shell/output"
)

//
// Create ...
//
type Create struct {
	Ignore bool
	Order  uint64
	Type   string
}

//
// GetCreateTag ...
//
func GetCreateTag(tag reflect.StructTag) *Create {
	var err error
	tagValue, ok := tag.Lookup("create")

	if !ok {
		return nil
	}

	ret := &Create{
		Ignore: false,
		Type:   "value",
	}

	args := strings.Split(tagValue, ",")

	for _, a := range args {
		switch {
		case (a == "ignore"):
			ret.Ignore = true

		case (strings.HasPrefix(a, "order=")):
			ret.Order, err = strconv.ParseUint(a[6:], 10, 32)

			if err != nil {
				output.Println(err)
				return nil
			}

		case (strings.HasPrefix(a, "type=")):
			ret.Type = a[5:]
		}
	}

	return ret
}
