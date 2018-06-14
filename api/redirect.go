package api

import (
	"net/http"
	"time"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

//
// RedirectList ...
//
func (a *myraAPI) RedirectList(domain string, filter string) ([]vo.RedirectVO, error) {
	ret, err := a.request(
		http.MethodGet,
		"/redirects/"+domain+"/1?search="+filter,
		nil,
	)

	if err != nil {
		return []vo.RedirectVO{}, err
	}

	queryVO := struct {
		vo.QueryVO

		List []vo.RedirectVO `json:"list"`
	}{}

	err = a.unmarshalResponse(ret, &queryVO)
	if err != nil {
		return []vo.RedirectVO{}, err
	}

	return queryVO.List, nil
}

//
// saveRedirect ...
//
func (a *myraAPI) saveRedirect(domain string, obj vo.RedirectVO) error {
	method := http.MethodPut
	if obj.BaseEntityVO.ID != nil {
		method = http.MethodPost
	}

	ret, err := a.request(method, "/redirects/"+domain, obj)

	if err != nil {
		return err
	}

	resultVO := struct {
		vo.ResultVO

		TargetObject []vo.RedirectVO `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}

//
// RemoveRedirect ...
//
func (a *myraAPI) removeRedirect(domain string, obj vo.RedirectVO) error {
	method := http.MethodDelete

	ret, err := a.request(method, "/redirects/"+domain, struct {
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

		TargetObject []vo.RedirectVO `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}
