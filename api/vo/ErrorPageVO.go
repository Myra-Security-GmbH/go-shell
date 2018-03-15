package vo

import (
	validator "gopkg.in/go-playground/validator.v9"
)

//
// ErrorPageVOList ...
//
type ErrorPageVOList struct {
	SubDomainName string                 `json:"subDomainName"`
	Error         map[string]ErrorPageVO `json:"error"`
}

//
// ErrorPageVO ...
//
type ErrorPageVO struct {
	BaseEntityVO `create:"ignore"`

	ErrorCode int    `json:"errorCode" create:"order=0"`
	Content   string `json:"content" create:"order=1,type=file"`
}

//
// Validate ...
//
func (vo ErrorPageVO) Validate(sl validator.StructLevel) {
}

//
// ResetDatabaseState ...
//
func (vo ErrorPageVO) ResetDatabaseState() interface{} {
	vo.BaseEntityVO.ID = nil
	vo.BaseEntityVO.Created = nil
	vo.BaseEntityVO.Modified = nil

	return vo
}
