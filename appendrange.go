package piscine

func AppendRange(min, max int) []int {
	var resultat []int
	if min > max {
		return nil
	}
	for i := min; i < max; i++ {
		resultat = append(resultat, i)
	}
	return resultat
}
