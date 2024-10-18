package piscine

func Join(strs []string, sep string) string {
	var result string
	n := 0
	for i := range strs {
		n = i + 1
	}
	for i, str := range strs {
		if i == n-1 {
			result += str
		} else {
			result += str + sep
		}
	}
	return result
}
