package listing

import (
	"strconv"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

type redirectRow struct {
	data vo.RedirectVO
}

func (r *redirectRow) GetID() uint64 {
	return *r.data.ID
}

func (r *redirectRow) FormatFlags() string {
	return "rw-"
}

func (r *redirectRow) IsAvailableForContextSwitch() bool {
	return false
}

func (r *redirectRow) GetName() string {
	return r.data.Source
}

func (r *redirectRow) GetEntity() interface{} {
	return r.data
}

func (r *redirectRow) GetColumns(long bool, verbose bool) []interface{} {
	ret := []interface{}{
		strconv.FormatUint(*r.data.ID, 10),
		r.data.Modified.ToUnixDate(),
	}

	if long {
		ret = append(ret, r.data.Type)
		ret = append(ret, r.data.Destination)
	}

	return ret
}
