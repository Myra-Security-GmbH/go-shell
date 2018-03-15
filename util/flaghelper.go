package util

//
// FlagOut ...
//
func FlagOut(flag string, show bool) string {
	if show {
		return flag
	}

	return "-"
}
