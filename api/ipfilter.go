package api

import (
	"net/http"
	"time"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

//
// IPFilterList ...
//
func (a *myraAPI) IPFilterList(domain string, search string) ([]vo.IPFilterVO, error) {
	ret, err := a.request(
		http.MethodGet,
		"/ipfilter/"+domain+"/1?search="+search,
		nil,
	)

	if err != nil {
		return []vo.IPFilterVO{}, err
	}

	queryVO := struct {
		vo.QueryVO

		List []vo.IPFilterVO `json:"list"`
	}{}

	err = a.unmarshalResponse(ret, &queryVO)

	if err != nil {
		return []vo.IPFilterVO{}, err
	}

	return queryVO.List, nil
}

//
// saveIPFilter ...
//
func (a *myraAPI) saveIPFilter(domain string, obj vo.IPFilterVO) error {
	method := http.MethodPut
	if obj.BaseEntityVO.ID != nil {
		method = http.MethodPost
	}

	ret, err := a.request(method, "/ipfilter/"+domain, obj)

	if err != nil {
		return err
	}

	resultVO := struct {
		vo.ResultVO

		TargetObject []vo.IPFilterVO `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}

//
// RemoveIPFilter ...
//
func (a *myraAPI) removeIPFilter(domain string, obj vo.IPFilterVO) error {
	method := http.MethodDelete

	ret, err := a.request(method, "/ipfilter/"+domain, struct {
		ID       uint64 `json:"id"`
		Modified string `json:"modified"`
	}{
		ID:       *obj.BaseEntityVO.ID,
		Modified: (*obj.BaseEntityVO.Modified).Format(time.RFC3339),
	})

	if err != nil {
		return err
	}

	resultVO := struct {
		vo.ResultVO

		TargetObject []vo.IPFilterVO `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}
