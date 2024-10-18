package piscine

func IsPrintable(str string) bool {
	x := []rune(str)

	for i := 0; i <= len(x)-1; i++ {
		if (x[i] >= 0) && (x[i] <= 31) {
			return false
		}
	}
	return true
}
