package api

import (
	"net/http"
	"time"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

//
// SslCertList ...
//
func (a *myraAPI) SslCertList(domain string, search string) ([]vo.SslCertVO, error) {
	ret, err := a.request(
		http.MethodGet,
		"/certificates/"+domain+"/1?search="+search,
		nil,
	)

	if err != nil {
		return []vo.SslCertVO{}, err
	}

	queryVO := struct {
		vo.QueryVO

		List []vo.SslCertVO `json:"list"`
	}{}

	err = a.unmarshalResponse(ret, &queryVO)

	if err != nil {
		return []vo.SslCertVO{}, err
	}

	return queryVO.List, nil
}

//
// saveIPFilter ...
//
func (a *myraAPI) saveSslCert(domain string, obj vo.SslCertVO) error {
	method := http.MethodPut
	if obj.BaseEntityVO.ID != nil {
		method = http.MethodPost
	}

	ret, err := a.request(method, "/certificates/"+domain, obj)

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
func (a *myraAPI) removeSslCert(domain string, obj vo.SslCertVO) error {
	method := http.MethodDelete

	ret, err := a.request(method, "/certificates/"+domain, struct {
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

		TargetObject []vo.SslCertVO `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}
