package vo

import (
	validator "gopkg.in/go-playground/validator.v9"
)

//
// IPFilterVO ...
//
type IPFilterVO struct {
	BaseEntityVO `create:"ignore"`

	Type  string `json:"type" create:"order=1" validate:"required,eq=BLACKLIST|eq=WHITELIST"`
	Value string `json:"value" create:"order=0" validate:"required,cidr"`
}

// //
// // RowListing ...
// //
// func (vo IPFilterVO) RowListing(cmd *command.Command) []*listing.Row {
// 	lr := listing.NewRow(vo.Value)
// 	lr.Readable = true
//
// 	lr.Cols = []interface{}{
// 		*vo.BaseEntityVO.ID,
// 		vo.BaseEntityVO.Modified.ToUnixDate(),
// 		vo.Type,
// 	}
//
// 	return []*listing.Row{lr}
// }

//
// ResetDatabaseState ...
//
func (vo IPFilterVO) ResetDatabaseState() interface{} {
	vo.BaseEntityVO.ID = nil
	vo.BaseEntityVO.Created = nil
	vo.BaseEntityVO.Modified = nil

	return vo
}

//
// Validate ...
//
func (vo IPFilterVO) Validate(sl validator.StructLevel) {
}
