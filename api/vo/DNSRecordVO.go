package vo

import (
	validator "gopkg.in/go-playground/validator.v9"
)

//
// DNSRecordVO ...
//
type DNSRecordVO struct {
	BaseEntityVO `create:"ignore"`

	Name             string `json:"name" validate:"required" create:"order=0"`
	TTL              uint   `json:"ttl" validate:"required,gte=300,lte=86400" create:"order=3"`
	Type             string `json:"recordType" validate:"required,eq=A|eq=AAAA|eq=CNAME|eq=MX|eq=SRV|eq=PTR|eq=NS|eq=TXT" create:"order=1"`
	AlternativeCName string `json:"alternativeCname" create:"ignore"`
	Value            string `json:"value" validate:"required" create:"order=2"`
	Priority         int    `json:"priority" validate:"gte=0,lte=65535"`
	Port             uint   `json:"port" validate:"gte=0,lte=65535"`
	Protected        bool   `json:"active" create:"order=4"`
	Enabled          bool   `json:"enabled"`
}

//
// Validate ...
//
func (vo DNSRecordVO) Validate(sl validator.StructLevel) {
	v := sl.Current().Interface().(DNSRecordVO)

	if !v.CanBeProtected() && v.Protected {
		sl.ReportError(v.Protected, "Protected", "protected", "protected", "")
	}

	switch v.Type {
	case RecordTypeA:
		sl.Validator().Var(v.Value, "ipv4")
		break

	case RecordTypeAAAA:
		sl.Validator().Var(v.Value, "ipv6")
		break

	case RecordTypeCNAME:
		sl.Validator().Var(v.Value, "fqdn")
		break

	case RecordTypeMX:
		sl.Validator().Var(v.Value, "fqdn")
		break

	case RecordTypeSRV:
		sl.Validator().Var(v.Value, "fqdn")
		break

	case RecordTypePTR:
		sl.Validator().Var(v.Value, "fqdn")
		break

	case RecordTypeNS:
		sl.Validator().Var(v.Value, "fqdn")
		break

		// case RecordTypeTXT:
		// 	break
	}
}

//
// CanBeProtected ...
//
func (vo DNSRecordVO) CanBeProtected() bool {
	return vo.Type == RecordTypeA ||
		vo.Type == RecordTypeAAAA ||
		vo.Type == RecordTypeCNAME
}

//
// ResetDatabaseState ...
//
func (vo DNSRecordVO) ResetDatabaseState() interface{} {
	vo.BaseEntityVO.ID = nil
	vo.BaseEntityVO.Created = nil
	vo.BaseEntityVO.Modified = nil

	return vo
}
