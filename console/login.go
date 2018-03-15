package console

//
// Login ...
//
type Login struct {
	Name   string
	APIKey string
	Secret string
}

// //
// // RowListing ...
// //
// func (l Login) RowListing(cmd *command.Command) []*listing.Row {
// 	lr := listing.NewRow(l.Name)
// 	lr.SwitchContext = true
//
// 	_, long := cmd.Flags["l"]
//
// 	lr.Cols = []interface{}{}
//
// 	if long {
// 		lr.Cols = append(lr.Cols, l.APIKey)
// 	}
//
// 	return []*listing.Row{lr}
// }
