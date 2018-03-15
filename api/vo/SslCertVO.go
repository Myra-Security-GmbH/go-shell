package vo

import (
	validator "gopkg.in/go-playground/validator.v9"
)

//
// SslCertVO ...
//
type SslCertVO struct {
	BaseEntityVO `create:"ignore"`

	SubjectAlternatives []string            `json:"subjectAlternatives"`
	Key                 string              `json:"key" validate:"required"`
	Wildcard            bool                `json:"wildcard" `
	ExtendedValidation  bool                `json:"extendedValidation"`
	IPList              []string            `json:"ipList"`
	Subdomains          []string            `json:"subDomains"`
	CertToRefresh       int                 `json:"certToRefresh"`
	CertRefreshForced   bool                `json:"certRefreshForced"`
	SniAllowed          bool                `json:"sniAllowed"`
	Password            string              `json:"password"`
	Subject             string              `json:"subject" validate:"required"`
	Algorithm           string              `json:"algorithm"`
	ValidFrom           string              `json:"validFrom"`
	ValidTo             string              `json:"validTo"`
	Cert                string              `json:"cert"`
	Fingerprint         string              `json:"fingerprint"`
	SerialNumber        string              `json:"serialNumber"`
	Intermediates       []SslIntermediateVO `json:"intermediates"`
}

//
// Validate ...
//
func (vo SslCertVO) Validate(sl validator.StructLevel) {
	//v := sl.Current().Interface().(SslCertVO)
}

//
// ResetDatabaseState ...
//
func (vo SslCertVO) ResetDatabaseState() interface{} {
	vo.BaseEntityVO.ID = nil
	vo.BaseEntityVO.Created = nil
	vo.BaseEntityVO.Modified = nil

	return vo
}
