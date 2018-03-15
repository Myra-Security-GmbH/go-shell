package listing

import (
	"strconv"

	"myracloud.com/myra-shell/api/vo"
)

type sslrow struct {
	data vo.SslCertVO
}

func (r *sslrow) GetID() uint64 {
	return *r.data.ID
}

func (r *sslrow) FormatFlags() string {
	return "rw-"
}

func (r *sslrow) IsAvailableForContextSwitch() bool {
	return false
}

func (r *sslrow) GetName() string {
	return r.data.Subject
}

func (r *sslrow) GetEntity() interface{} {
	return r.data
}

func (r *sslrow) GetColumns(long bool, verbose bool) []interface{} {
	ret := []interface{}{}

	ret = append(ret, strconv.FormatUint(*r.data.ID, 10))
	ret = append(ret, r.data.Modified.ToUnixDate())
	ret = append(ret, strconv.FormatInt(int64(len(r.data.Subdomains)), 10))

	return ret
}
