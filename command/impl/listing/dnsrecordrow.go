package listing

import (
	"strconv"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
	"github.com/Myra-Security-GmbH/myra-shell/util"
)

type dnsrecordrow struct {
	data vo.DNSRecordVO
}

func (r *dnsrecordrow) GetID() uint64 {
	return *r.data.ID
}

func (r *dnsrecordrow) FormatFlags() string {
	return "rwx" +
		util.FlagOut("e", r.data.Enabled) +
		util.FlagOut("p", r.data.Protected)
}

func (r *dnsrecordrow) IsAvailableForContextSwitch() bool {
	return r.data.CanBeProtected()
}

func (r *dnsrecordrow) GetName() string {
	return r.data.Name
}

func (r *dnsrecordrow) GetEntity() interface{} {
	return r.data
}

func (r *dnsrecordrow) GetColumns(long bool, verbose bool) []interface{} {
	ret := []interface{}{}

	ret = append(ret, strconv.FormatUint(*r.data.ID, 10))
	ret = append(ret, r.data.Modified.ToUnixDate())

	if long {
		ret = append(ret, r.data.Type)
	}

	return ret
}
