package listing

import (
	"strconv"
)

//
// Row ...
//
type Row interface {
	GetName() string
	GetColumns(long bool, verbose bool) []interface{}
	GetID() uint64
	GetEntity() interface{}
	FormatFlags() string
	IsAvailableForContextSwitch() bool
}

//
// RowListing ...
//
type RowListing []Row

func (rl RowListing) buildFormatingTemplate(long bool, verbose bool) string {
	var ret string
	var columnLen []int64

	for _, r := range rl {
		for idx, c := range r.GetColumns(long, verbose) {
			if len(columnLen) > idx {
				columnLen[idx] = maxLen(int64(len(c.(string))), columnLen[idx])
			} else {
				columnLen = append(columnLen, int64(len(c.(string))))
			}
		}
	}

	for _, cl := range columnLen {
		ret += "%" + strconv.FormatInt(cl, 10) + "s "
	}

	return ret
}
