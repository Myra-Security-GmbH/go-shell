package listing

func formatFlag(tVal rune, fVal rune, set bool) string {
	if set {
		return string(tVal)
	}

	return string(fVal)
}

func maxLen(a int64, b int64) int64 {
	if a > b {
		return a
	}

	return b
}
