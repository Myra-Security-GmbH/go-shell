package util

//
// HasOption checks if one of the given option elements is set.
//
func HasOption(options map[string]bool, option ...string) bool {
	for _, o := range option {
		_, ok := options[o]

		if ok {
			return true
		}
	}

	return false
}
