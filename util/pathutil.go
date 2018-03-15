package util

import "os"

//
// GetFirstPathMatches returns the first path
// that is available.
//
func GetFirstPathMatches(path []string) string {
	for _, p := range path {
		_, err := os.Stat(p)

		if err == nil {
			return p
		}
	}

	return ""
}
