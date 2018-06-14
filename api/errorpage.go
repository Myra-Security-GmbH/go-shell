package api

import (
	"net/http"
	"time"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

//
// ErrorPage ...
//
func (a *myraAPI) ErrorPage(domain string, search string) ([]vo.ErrorPageVOList, error) {
	ret, err := a.request(
		http.MethodGet,
		"/errorpages/"+domain+"/1?search?"+search,
		nil,
	)

	if err != nil {
		return []vo.ErrorPageVOList{}, err
	}

	queryVO := struct {
		vo.QueryVO

		List []vo.ErrorPageVOList `json:"list"`
	}{}

	err = a.unmarshalResponse(ret, &queryVO)

	if err != nil {
		return []vo.ErrorPageVOList{}, err
	}

	return queryVO.List, nil
}

//
// saveIPFilter ...
//
func (a *myraAPI) saveErrorPage(domain string, obj vo.ErrorPageVO) error {
	method := http.MethodPut
	if obj.BaseEntityVO.ID != nil {
		method = http.MethodPost
	}

	ret, err := a.request(method, "/errorpages/"+domain, obj)

	if err != nil {
		return err
	}

	resultVO := struct {
		vo.ResultVO

		TargetObject []vo.ErrorPageVOList `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}

//
// RemoveIPFilter ...
//
func (a *myraAPI) removeErrorPage(domain string, obj vo.ErrorPageVO) error {
	method := http.MethodDelete

	ret, err := a.request(method, "/errorpages/"+domain, struct {
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

		TargetObject vo.ErrorPageVOList `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}
