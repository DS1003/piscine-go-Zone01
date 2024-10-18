package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	choumi := os.Args
	for _, w := range choumi[0] {
		if w != '.' && w != '/' {
			z01.PrintRune(w)
		}
	}
	z01.PrintRune('\n')
}
