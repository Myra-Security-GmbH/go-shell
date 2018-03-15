package listing

import (
	"strconv"

	"myracloud.com/myra-shell/api/vo"
)

type domainRow struct {
	data vo.DomainVO
}

func (r *domainRow) GetID() uint64 {
	return *r.data.ID
}

func (r *domainRow) FormatFlags() string {
	return "--x"
}

func (r *domainRow) IsAvailableForContextSwitch() bool {
	return true
}

func (r *domainRow) GetName() string {
	return r.data.Name
}

func (r *domainRow) GetEntity() interface{} {
	return r.data
}

func (r *domainRow) GetColumns(long bool, verbose bool) []interface{} {
	return []interface{}{
		strconv.FormatUint(*r.data.ID, 10),
		r.data.Modified.ToUnixDate(),
	}
}
