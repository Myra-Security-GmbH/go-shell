package listing

import (
	"strconv"

	"github.com/Myra-Security-GmbH/myra-shell/api/types"
	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

type errorpageRow struct {
	subDomainName string
	data          vo.ErrorPageVO
}

func (r *errorpageRow) GetID() uint64 {
	return 0
}

func (r *errorpageRow) FormatFlags() string {
	return "rw-"
}

func (r *errorpageRow) IsAvailableForContextSwitch() bool {
	return false
}

func (r *errorpageRow) GetName() string {
	return r.subDomainName
}

func (r *errorpageRow) GetEntity() interface{} {
	return nil
}

func (r *errorpageRow) GetColumns(long bool, verbose bool) []interface{} {
	return []interface{}{
		strconv.FormatInt(int64(len(r.data.Content)), 10),
		types.DateTimeNow().ToUnixDate(),
		strconv.FormatInt(int64(r.data.ErrorCode), 10),
	}
}
