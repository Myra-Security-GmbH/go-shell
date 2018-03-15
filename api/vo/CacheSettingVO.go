package vo

import (
	validator "gopkg.in/go-playground/validator.v9"
)

//
// CacheSettingVO ...
//
type CacheSettingVO struct {
	BaseEntityVO `create:"ignore"`

	Path string `json:"path" create:"order=0" validate:"required"`
	TTL  string `json:"ttl" create:"order=2" validate:"required,numeric"`
	Type string `json:"type" create:"order=1" validate:"required,eq=prefix|eq=suffix|eq=exact"`
}

// //
// // RowListing ...
// //
// func (c CacheSettingVO) RowListing(cmd *command.Command) []*listing.Row {
// 	_, exist := cmd.Flags["l"]
//
// 	lr := listing.NewRow(c.Path)
//
// 	lr.Cols = []interface{}{
// 		*c.BaseEntityVO.ID,
// 		c.Modified.ToUnixDate(),
// 		c.Type,
// 	}
//
// 	if exist {
// 		lr.Cols = append(lr.Cols, fmt.Sprintf("ttl=\"%s\"", c.TTL))
// 		lr.Cols = append(
// 			lr.Cols,
// 			fmt.Sprintf("created=\"%s\"", c.Created.ToUnixDate()),
// 		)
// 	}
//
// 	return []*listing.Row{lr}
// }

//
// ResetDatabaseState ...
//
func (c CacheSettingVO) ResetDatabaseState() interface{} {
	c.BaseEntityVO.ID = nil
	c.BaseEntityVO.Created = nil
	c.BaseEntityVO.Modified = nil

	return c
}

//
// Validate ...
//
func (c CacheSettingVO) Validate(sl validator.StructLevel) {
}
