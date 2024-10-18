package piscine

import "github.com/01-edu/z01"

func main() {
	PrintNbrInOrder(32132)
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var array []int
		eachValue := 0

		countary := 0
		var minval int
		for n != 0 {
			eachValue = n % 10
			n /= 10
			array = append(array, eachValue)
		}

		for count := range array {
			countary = count + 1
		}
		for i := 0; i < countary-1; i++ {
			for j := 0; j < countary-i-1; j++ {
				if array[j] > array[j+1] {
					minval = array[j]
					array[j] = array[j+1]
					array[j+1] = minval
				}
			}
		}
		for i := 0; i < countary; i++ {
			z01.PrintRune(rune(array[i] + 48))
		}
	}
}
