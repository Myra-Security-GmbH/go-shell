package api

import (
	"net/http"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

//
// DomainList ...
//
func (a *myraAPI) DomainList(search string) ([]vo.DomainVO, error) {
	ret, err := a.request(http.MethodGet, "/domains/1?search="+search, nil)

	if err != nil {
		return []vo.DomainVO{}, err
	}

	queryVO := struct {
		vo.QueryVO

		List []vo.DomainVO `json:"list"`
	}{}

	err = a.unmarshalResponse(ret, &queryVO)

	if err != nil {
		return []vo.DomainVO{}, err
	}

	return queryVO.List, nil
}

//
// DomainByName ...
//
func (a *myraAPI) DomainByName(name string) (*vo.DomainVO, error) {
	list, _ := a.DomainList(name)

	if len(list) == 1 {
		return &list[0], nil
	}

	return nil, nil
}
