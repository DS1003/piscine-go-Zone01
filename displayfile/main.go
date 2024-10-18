package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) >= 2 {
		fmt.Println("Too many arguments")
	} else {
		file, er := ioutil.ReadFile(args[0])
		if er != nil {
			fmt.Println(er)
		}
		fmt.Print(string(file))
	}
}
