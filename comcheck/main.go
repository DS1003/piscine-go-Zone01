package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	find := []string{"01", "galaxy", "galaxy 01"}
	for _, i := range arg {
		for _, j := range find {
			if i == j {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
