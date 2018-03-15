package vo

import (
	validator "gopkg.in/go-playground/validator.v9"
)

//
// SslIntermediateVO ...
//
type SslIntermediateVO struct {
	BaseEntityVO `create:"ignore"`

	Subject      string `json:"subject" validate:"required"`
	Algorithm    string `json:"algorithm"`
	ValidFrom    string `json:"validFrom"`
	ValidTo      string `json:"validTo"`
	Cert         string `json:"cert"`
	Fingerprint  string `json:"fingerprint"`
	SerialNumber string `json:"serialNumber"`
	Issuer       string `json:"issuer"`
}

//
// Validate ...
//
func (vo SslIntermediateVO) Validate(sl validator.StructLevel) {
	//v := sl.Current().Interface().(SslCertVO)
}

//
// ResetDatabaseState ...
//
func (vo SslIntermediateVO) ResetDatabaseState() interface{} {
	vo.BaseEntityVO.ID = nil
	vo.BaseEntityVO.Created = nil
	vo.BaseEntityVO.Modified = nil

	return vo
}
