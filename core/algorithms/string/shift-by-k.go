package string

func shiftByK(s string, k int) string {
	k = k % len(s)
	if k == 0 {
		return s
	}

	return s[len(s)-k:] + s[:len(s)-k]
}
