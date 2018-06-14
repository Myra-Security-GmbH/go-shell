package listing

import (
	"strconv"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

type cacheSettingsRow struct {
	data vo.CacheSettingVO
}

func (r *cacheSettingsRow) GetID() uint64 {
	return *r.data.ID
}

func (r *cacheSettingsRow) FormatFlags() string {
	return "rw-"
}

func (r *cacheSettingsRow) IsAvailableForContextSwitch() bool {
	return false
}

func (r *cacheSettingsRow) GetName() string {
	return r.data.Path
}

func (r *cacheSettingsRow) GetEntity() interface{} {
	return r.data
}

func (r *cacheSettingsRow) GetColumns(long bool, verbose bool) []interface{} {
	ret := []interface{}{
		strconv.FormatUint(*r.data.ID, 10),
		r.data.Modified.ToUnixDate(),
	}

	if long {
		ret = append(ret, r.data.TTL)
		ret = append(ret, r.data.Type)
	}

	return ret
}
