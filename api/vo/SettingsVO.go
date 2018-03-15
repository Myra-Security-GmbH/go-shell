package vo

import (
	"fmt"
	"strconv"
)

// SettingsVO ...
type SettingsVO map[string]interface{}

// //
// // RowListing ...
// //
// func (vo SettingsVO) RowListing(cmd *command.Command) []*listing.Row {
// 	ret := []*listing.Row{}
// 	max := 0
//
// 	sortedList := make([]string, len(vo))
//
// 	for k := range vo {
// 		if strings.HasPrefix(k, "nginx_") || k == "proxy_cache_key" {
// 			continue
// 		}
//
// 		if len(k) > max {
// 			max = len(k)
// 		}
//
// 		sortedList = append(sortedList, k)
// 	}
//
// 	sort.Strings(sortedList)
//
// 	template := fmt.Sprintf("%%-%ds => %%s", max)
//
// 	for _, k := range sortedList {
// 		v, exists := vo[k]
//
// 		if !exists {
// 			continue
// 		}
//
// 		v = formatSettingsValue(v)
//
// 		lr := listing.NewRow(fmt.Sprintf(template, k, v))
// 		lr.Cols = []interface{}{
// 			0,
// 			listing.DateTimeNow().ToUnixDate(),
// 		}
//
// 		ret = append(ret, lr)
// 	}
//
// 	return ret
// }

//
//
//
func formatSettingsValue(v interface{}) string {
	var ret string

	switch val := v.(type) {
	case string:
		ret = fmt.Sprintf("\"%s\"", val)

	case int:
		ret = strconv.FormatInt(v.(int64), 10)

	case nil:
		ret = "null"

	case float64:
		ret = fmt.Sprintf("%.2f", val)

	case bool:
		ret = strconv.FormatBool(val)
	}

	return ret
}
