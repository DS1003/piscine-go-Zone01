package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	arrayS := []rune(s)
	for i := 0; i < len(arrayS); i++ {
		z01.PrintRune(arrayS[i])
	}
	z01.PrintRune('\n')
}

func main() {
	lengthOfArg := len(os.Args) - 1
	if lengthOfArg%2 == 0 {
		EvenMsg := "I have an even number of arguments"
		printStr(EvenMsg)
	} else {
		OddMsg := "I have an odd number of arguments"
		printStr(OddMsg)
	}
}
