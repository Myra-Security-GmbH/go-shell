package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"myracloud.com/myra-shell/api/vo"
)

func (a *myraAPI) request(method string, url string, obj interface{}) (*http.Response, error) {
	var content []byte
	var err error

	if obj != nil {
		content, err = json.Marshal(&obj)

		if err != nil {
			return nil, err
		}

		//fmt.Println(string(content))
	} else {
		content = []byte("")
	}

	request, err := http.NewRequest(
		method,
		a.endpoint+"/"+a.language+"/rapi"+url,
		bytes.NewReader(content),
	)

	if err != nil {
		return nil, err
	}

	t := time.Now().Format(time.RFC3339)

	sig, err := a.signature.Build(
		string(content),
		method,
		"/"+a.language+"/rapi"+url,
		"application/json",
		t,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "MYRA "+a.apiKey+":"+sig)
	request.Header.Add("Date", t)
	request.Header.Add("Content-Type", "application/json")

	ret, err := a.client.Do(request)

	if err != nil {
		return nil, err
	}

	if ret.StatusCode == 403 {
		return nil, errors.New("Permission denied")
	}

	return ret, nil
}

func (a *myraAPI) unmarshalResponse(res *http.Response, data interface{}) error {
	tmp, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if res.StatusCode < 200 && res.StatusCode > 300 {
		return fmt.Errorf("API returned status code: %d\n%s", res.StatusCode, tmp)
	}

	err = json.Unmarshal(tmp, data)

	if err != nil {
		return err
	}

	return a.buildErrorMessage(data)
}

//
// buildErrorMessage generates an error message out of the ViolationList.
//
func (a *myraAPI) buildErrorMessage(data interface{}) error {
	t := reflect.ValueOf(data).Elem()

	if t.Kind() != reflect.Struct || !t.FieldByName("Error").Bool() {
		return nil
	}

	if !t.FieldByName("ViolationList").IsValid() {
		return nil
	}

	msg := ""
	list := t.FieldByName("ViolationList").Interface()
	val := reflect.ValueOf(list)

	if val.Kind() != reflect.Slice {
		return nil
	}

	for i := 0; i < val.Len(); i++ {
		k := val.Index(i).Interface().(vo.ViolationVO)

		msg += "- "

		if k.Path != "" {
			msg += k.Path + ": "
		}

		msg += k.Message + "\n"
	}

	return errors.New(msg)
}

//
// SaveEntity ...
//
func (a *myraAPI) SaveEntity(identifer string, entity interface{}) error {
	r := reflect.ValueOf(entity)

	if r.Kind() == reflect.Ptr {
		entity = r.Elem().Interface()
	}

	err := a.validator.Struct(entity)

	if err != nil {
		return err
	}

	switch v := entity.(type) {
	case vo.IPFilterVO:
		return a.saveIPFilter(identifer, v)

	case vo.CacheSettingVO:
		return a.saveCacheSetting(identifer, v)

	case vo.RedirectVO:
		return a.saveRedirect(identifer, v)

	case vo.DNSRecordVO:
		return a.saveDNSRecord(identifer, v)
	}

	return errors.New("Could not handle entity type")
}

//
// RemoveEntity ...
//
func (a *myraAPI) RemoveEntity(identifier string, entity interface{}) error {
	r := reflect.ValueOf(entity)

	if r.Kind() == reflect.Ptr {
		entity = r.Elem().Interface()
	}

	switch v := entity.(type) {
	case vo.IPFilterVO:
		return a.removeIPFilter(identifier, v)

	case vo.CacheSettingVO:
		return a.removeCacheSetting(identifier, v)

	case vo.RedirectVO:
		return a.removeRedirect(identifier, v)

	case vo.DNSRecordVO:
		return a.removeDNSRecord(identifier, v)
	}

	return errors.New("Could not handle entity type")
}
