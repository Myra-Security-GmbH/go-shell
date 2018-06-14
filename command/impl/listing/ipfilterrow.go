package listing

import (
	"strconv"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

type ipfilterrow struct {
	data vo.IPFilterVO
}

func (r *ipfilterrow) GetID() uint64 {
	return *r.data.ID
}

func (r *ipfilterrow) FormatFlags() string {
	return "rw-"
}

func (r *ipfilterrow) IsAvailableForContextSwitch() bool {
	return false
}

func (r *ipfilterrow) GetName() string {
	return r.data.Value
}

func (r *ipfilterrow) GetEntity() interface{} {
	return r.data
}

func (r *ipfilterrow) GetColumns(long bool, verbose bool) []interface{} {
	ret := []interface{}{}

	ret = append(ret, strconv.FormatUint(*r.data.ID, 10))
	ret = append(ret, r.data.Modified.ToUnixDate())
	ret = append(ret, r.data.Type)

	return ret
}
