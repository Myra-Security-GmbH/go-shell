package vo

import "myracloud.com/myra-shell/api/types"

//
// DomainVO ...
//
type DomainVO struct {
	BaseEntityVO `create:"ignore"`

	Name        string         `json:"name"`
	PausedUntil types.DateTime `json:"pausedUntil"`
	Paused      bool           `json:"paused"`
	Owned       bool           `json:"owned"`
	Maintenance bool           `json:"maintenance"`
	AutoUpdate  bool           `json:"autoUpdate"`
}

// //
// // RowListing ...
// //
// func (vo DomainVO) RowListing(cmd *command.Command) []*listing.Row {
// 	lr := listing.NewRow(vo.Name)
// 	lr.Readable = true
// 	lr.Writeable = vo.Owned
// 	lr.SwitchContext = true
//
// 	lr.Flags["maintenance"] = vo.Maintenance
// 	lr.Flags["paused"] = vo.Paused
// 	lr.Flags["autoUpdate"] = vo.AutoUpdate
//
// 	lr.Cols = []interface{}{
// 		*vo.BaseEntityVO.ID,
// 		vo.BaseEntityVO.Modified.ToUnixDate(),
// 	}
//
// 	return []*listing.Row{lr}
// }

//
// ResetDatabaseState ...
//
func (vo DomainVO) ResetDatabaseState() interface{} {
	vo.BaseEntityVO.ID = nil
	vo.BaseEntityVO.Created = nil
	vo.BaseEntityVO.Modified = nil

	return vo
}
