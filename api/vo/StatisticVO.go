package vo

import "myracloud.com/myra-shell/api/types"

//
// DatasourceVO ...
//
type DatasourceVO struct {
	Source string `json:"source"`
	Type   string `json:"type"`
}

//
// StatisticQueryVO ...
//
type StatisticQueryVO struct {
	AggregationInterval string                  `json:"aggregationInterval"`
	DataSource          map[string]DatasourceVO `json:"dataSources"`
	StartDate           *types.DateTime         `json:"startDate"`
	EndDate             *types.DateTime         `json:"endDate"`
	FQDN                []string                `json:"fqdn"`
	Type                string                  `json:"type"`
}

//
// StatisticResultVO ...
//
type StatisticResultVO map[string]interface{}

//
// GetHistogramValues ...
//
func (vo StatisticResultVO) GetHistogramValues(source string) []float64 {
	ret := []float64{}

	_, ok := vo[source]

	if !ok {
		return ret
	}

	switch v := vo[source].(type) {
	case map[string]interface{}:
		for _, val := range v {
			v := val.(map[string]interface{})

			ret = append(ret, v["value"].(float64))
		}
		break
	}

	return ret
}

// //
// // RowListing ...
// //
// func (vo StatisticResultVO) RowListing(cmd *command.Command) []*listing.Row {
// 	ret := []*listing.Row{}
// 	sortedList := []string{}
//
// 	maxLen := 0
// 	for name := range vo {
// 		if maxLen < len(name) {
// 			maxLen = len(name)
// 		}
//
// 		sortedList = append(sortedList, name)
// 	}
//
// 	sort.Strings(sortedList)
//
// 	template := fmt.Sprintf("%%-%ds => %%s", maxLen)
//
// 	for _, name := range sortedList {
// 		val := vo[name]
//
// 		switch val.(type) {
// 		case map[string]interface{}:
// 			e := val.(map[string]interface{})
//
// 			if e["sum"] != nil {
// 				var value string
// 				if strings.HasPrefix(name, "bytes") {
// 					value = output.AsBytes(uint64(e["sum"].(float64)), 3)
// 				} else {
// 					value = output.AsNumber(e["sum"].(float64), 0)
// 				}
//
// 				ret = append(
// 					ret,
// 					listing.NewRowEx(
// 						fmt.Sprintf(
// 							template,
// 							name,
// 							value,
// 						),
// 						true,
// 						false,
// 						false,
// 						[]interface{}{0, listing.DateTimeNow().ToUnixDate()},
// 					),
// 				)
// 			}
// 			break
// 		}
// 	}
//
// 	return ret
// }

//
// StatisticVO ...
//
type StatisticVO struct {
	Query  StatisticQueryVO  `json:"query"`
	Result StatisticResultVO `json:"result,omitempty"`
}

//
// NewDataSourceVO ...
//
func NewDataSourceVO(source string, typ string) DatasourceVO {
	return DatasourceVO{
		Source: source,
		Type:   typ,
	}
}
