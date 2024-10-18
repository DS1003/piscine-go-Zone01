package piscine

func Capitalize(s string) string {
	runary := []rune(s)
	for i, char := range runary {
		if isNumberOrAlph(char) {
			if i == 0 || isNumberOrAlph(runary[i-1]) == false {
				if runary[i] >= 'a' && runary[i] <= 'z' {
					runary[i] = char - 32
				}
			} else {
				if runary[i] >= 'A' && runary[i] <= 'Z' {
					runary[i] = char + 32
				}
			}
		}
	}
	return string(runary)
}

func isNumberOrAlph(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}
