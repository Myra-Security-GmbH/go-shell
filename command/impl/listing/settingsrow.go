package listing

import (
	"fmt"
	"strconv"

	"github.com/Myra-Security-GmbH/myra-shell/api/types"
)

type settingsRow struct {
	name  string
	value interface{}
}

func (r *settingsRow) GetID() uint64 {
	return 0
}

func (r *settingsRow) FormatFlags() string {
	return "rw-"
}

func (r *settingsRow) IsAvailableForContextSwitch() bool {
	return false
}

func (r *settingsRow) GetName() string {
	return r.formattedSettingsValue()
}

func (r *settingsRow) GetEntity() interface{} {
	return nil
}

func (r *settingsRow) GetColumns(long bool, verbose bool) []interface{} {
	return []interface{}{
		"0",
		types.DateTimeNow().ToUnixDate(),
		r.name,
	}
}

//
//
//
func (r *settingsRow) formattedSettingsValue() string {
	var ret string

	switch val := r.value.(type) {
	case string:
		ret = fmt.Sprintf("\"%s\"", val)

	case int:
		ret = strconv.FormatInt(int64(val), 10)

	case nil:
		ret = "null"

	case float64:
		ret = fmt.Sprintf("%.2f", val)

	case bool:
		ret = strconv.FormatBool(val)
	}

	return ret
}
