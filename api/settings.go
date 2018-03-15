package api

import (
	"net/http"

	"myracloud.com/myra-shell/api/vo"
)

//
// Settings ...
//
func (a *myraAPI) Settings(domain string) (vo.SettingsVO, error) {
	ret, err := a.request(
		http.MethodGet,
		"/subdomainSetting/"+domain+"?flat",
		nil,
	)

	if err != nil {
		return vo.SettingsVO{}, err
	}

	list := &vo.SettingsVO{}

	err = a.unmarshalResponse(ret, list)

	if err != nil {
		return vo.SettingsVO{}, err
	}

	return *list, nil
}
