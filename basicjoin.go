package piscine

func BasicJoin(strs []string) string {
	r := ""
	for _, str := range strs {
		r = r + str
	}
	return r
}
