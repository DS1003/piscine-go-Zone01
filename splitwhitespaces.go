package piscine

func SplitWhiteSpaces(s string) []string {
	spl := 0
	res := []string{}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if spl != i {
				res = append(res, s[spl:i])
			}
			spl = i + 1
		} else if i == len(s)-1 {
			res = append(res, s[spl:i+1])
			spl = i + 1
		}
	}
	return res
}
