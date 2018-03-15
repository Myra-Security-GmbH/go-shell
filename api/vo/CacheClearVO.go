package vo

//
// CacheClearVO ...
//
type CacheClearVO struct {
	FQDN      string `json:"fqdn"`
	Resource  string `json:"resource"`
	Recursive bool   `json:"recursive"`
}

// //
// // RowListing ...
// //
// func (vo CacheClearVO) RowListing(cmd *command.Command) []*listing.Row {
// 	lr := listing.NewRow(vo.FQDN)
//
// 	lr.Cols = []interface{}{}
//
// 	lr.Cols = append(lr.Cols, vo.Resource)
// 	lr.Cols = append(lr.Cols, strconv.FormatBool(vo.Recursive))
//
// 	return []*listing.Row{lr}
// }
