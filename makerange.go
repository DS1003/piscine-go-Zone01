package piscine

func MakeRange(min, max int) []int {
	if max > min {
		resultat := make([]int, max-min)
		i, j := min, 0
		for i < max {
			resultat[j] = i
			i++
			j++
		}
		return resultat
	}
	return nil
}
