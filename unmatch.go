package piscine

func Unmatch(ar []int) int {
	for _, resultat := range ar {
		n := 0
		for _, r := range ar {
			if r == resultat {
				n++
			}
		}
		if n == 1 || n%2 == 1 {
			return resultat
		}
	}
	return -1
}
