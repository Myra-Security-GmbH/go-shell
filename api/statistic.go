package api

import (
	"net/http"

	"github.com/Myra-Security-GmbH/myra-shell/api/types"
	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

//
// GetStatistic ...
//
func (a *myraAPI) GetStatistic(
	domain string,
	interval string,
	startDate *types.DateTime,
	endDate *types.DateTime,
	source []vo.DatasourceVO,
) (*vo.StatisticVO, error) {
	sourceMap := make(map[string]vo.DatasourceVO)

	for _, s := range source {
		sourceMap[s.Source] = s
	}

	queryVO := &vo.StatisticVO{
		Query: vo.StatisticQueryVO{
			AggregationInterval: interval,
			FQDN:                []string{domain},
			Type:                "fqdn",
			StartDate:           startDate,
			EndDate:             endDate,
			DataSource:          sourceMap,
		},
	}

	ret, err := a.request(
		http.MethodPost,
		"/statistic/query",
		queryVO,
	)

	if err != nil {
		return nil, err
	}

	err = a.unmarshalResponse(ret, &queryVO)

	if err != nil {
		return &vo.StatisticVO{}, err
	}

	return queryVO, nil
}
