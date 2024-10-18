package piscine

func StrRev(s string) string {
	y := []rune(s)

	for i, j := 0, len(y)-1; i < j; i, j = i+1, j-1 {
		y[i], y[j] = y[j], y[i]
	}

	return string(y)
}
