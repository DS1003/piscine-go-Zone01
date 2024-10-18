package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := '0'
	b := '9'
	var count int = 0
	for i := a; i <= b; i++ {
		for j := a; j <= b; j++ {
			for k := a; k <= b; k++ {
				if i < j && j < k {
					if count > 0 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					count = 1
				}
			}
		}
	}
	z01.PrintRune('\n')
}
