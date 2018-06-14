package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

//
// DNSRecordList ...
//
func (a *myraAPI) DNSRecordList(
	domain string,
	filter *string,
	types *[]string,
	activeOnly bool,
) ([]vo.DNSRecordVO, error) {
	params := ""

	if filter != nil {
		params += "&filter=" + *filter
	}

	if types != nil {
		params += "&recordTypes=" + strings.Join(*types, ",")
	}

	if activeOnly {
		params += "&activeOnly=true"
	}

	if params != "" {
		params = "?" + strings.Trim(params, "&")
	}

	ret, err := a.request(
		http.MethodGet,
		"/dnsRecords/"+domain+"/1"+params,
		nil,
	)

	if err != nil {
		return []vo.DNSRecordVO{}, err
	}

	queryVO := struct {
		vo.QueryVO

		List []vo.DNSRecordVO `json:"list"`
	}{}

	err = a.unmarshalResponse(ret, &queryVO)

	if err != nil {
		return []vo.DNSRecordVO{}, err
	}

	return queryVO.List, nil
}

//
// saveDNSRecord ...
//
func (a *myraAPI) saveDNSRecord(domain string, obj vo.DNSRecordVO) error {
	method := http.MethodPut
	if obj.BaseEntityVO.ID != nil {
		method = http.MethodPost
	}

	ret, err := a.request(method, "/dnsRecords/"+domain, obj)

	if err != nil {
		return err
	}

	resultVO := struct {
		vo.ResultVO

		TargetObject []vo.DNSRecordVO `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}

//
// RemoveDNSRecord ...
//
func (a *myraAPI) removeDNSRecord(domain string, obj vo.DNSRecordVO) error {
	method := http.MethodDelete

	ret, err := a.request(method, "/dnsRecords/"+domain, struct {
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

		TargetObject []vo.DNSRecordVO `json:"targetObject"`
	}{}

	return a.unmarshalResponse(ret, &resultVO)
}
