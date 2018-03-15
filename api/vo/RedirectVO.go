package vo

//
// RedirectVO ...
//
type RedirectVO struct {
	BaseEntityVO `create:"ignore"`

	Source       string `json:"source"`
	Destination  string `json:"destination"`
	Type         string `json:"type"`
	MatchingType string `json:"matchingType"`
	ExpertMode   bool   `json:"expertMode" create:"ignore"`
}

// //
// // RowListing ...
// //
// func (vo RedirectVO) RowListing(cmd *command.Command) []*listing.Row {
// 	lr := listing.NewRow(vo.Source + " => " + vo.Destination)
// 	lr.Readable = true
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
func (vo RedirectVO) ResetDatabaseState() interface{} {
	vo.BaseEntityVO.ID = nil
	vo.BaseEntityVO.Created = nil
	vo.BaseEntityVO.Modified = nil

	return vo
}
