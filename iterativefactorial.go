package piscine

func IterativeFactorial(nb int) int {
	if nb < 25 {
		if nb == 0 {
			return 1
		}
		if nb < 0 {
			return 0
		}
		x := 1
		for y := nb; y >= 1; y-- {
			x *= y
		}
		return x
	}
	return 0
}
