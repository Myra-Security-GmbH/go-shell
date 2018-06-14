package api

import (
	"net/http"
	"time"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

//
// CacheSettingList ...
//
func (a *myraAPI) CacheSettingList(domain string, search string) ([]vo.CacheSettingVO, error) {
	ret, err := a.request(
		http.MethodGet,
		"/cacheSettings/"+domain+"/1?search="+search,
		nil,
	)

	if err != nil {
		return []vo.CacheSettingVO{}, err
	}

	queryVO := struct {
		vo.QueryVO

		List []vo.CacheSettingVO `json:"list"`
	}{}

	err = a.unmarshalResponse(ret, &queryVO)

	if err != nil {
		return []vo.CacheSettingVO{}, err
	}

	return queryVO.List, nil
}

//
// RemoveCacheSetting ...
//
func (a *myraAPI) removeCacheSetting(domain string, obj vo.CacheSettingVO) error {
	method := http.MethodDelete

	ret, err := a.request(method, "/cacheSettings/"+domain, struct {
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

		TargetObject []vo.CacheSettingVO `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}

//
// saveCacheSetting ...
//
func (a *myraAPI) saveCacheSetting(domain string, obj vo.CacheSettingVO) error {
	method := http.MethodPut
	if obj.BaseEntityVO.ID != nil {
		method = http.MethodPost
	}

	ret, err := a.request(method, "/cacheSettings/"+domain, obj)

	if err != nil {
		return err
	}

	resultVO := struct {
		vo.ResultVO

		TargetObject []vo.CacheSettingVO `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}

//
// CacheClear ...
//
func (a *myraAPI) CacheClear(domain string, fqdn string, pattern string, recursive bool) ([]vo.CacheClearVO, error) {
	method := http.MethodPut

	ret, err := a.request(method, "/cacheClear/"+domain, &vo.CacheClearVO{
		FQDN:      fqdn,
		Resource:  pattern,
		Recursive: recursive,
	})

	if err != nil {
		return nil, err
	}

	resultVO := struct {
		vo.ResultVO

		TargetObject []vo.CacheClearVO `json:"targetObject"`
	}{}

	err = a.unmarshalResponse(ret, &resultVO)

	if err != nil {
		return nil, err
	}

	return resultVO.TargetObject, nil
}
