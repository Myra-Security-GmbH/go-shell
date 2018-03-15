package listing

import "myracloud.com/myra-shell/console"

type loginRow struct {
	data console.Login
}

func (r *loginRow) GetID() uint64 {
	return 0
}

func (r *loginRow) FormatFlags() string {
	return "--x"
}

func (r *loginRow) IsAvailableForContextSwitch() bool {
	return true
}

func (r *loginRow) GetName() string {
	return r.data.Name
}

func (r *loginRow) GetEntity() interface{} {
	return nil
}

func (r *loginRow) GetColumns(long bool, verbose bool) []interface{} {
	if long {
		return []interface{}{
			r.data.APIKey,
		}

	}
	return []interface{}{}
}
