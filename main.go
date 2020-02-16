package main

import (
	"fmt"
	"strings"
)

func plus(a, b int) int {
	return a + b
}

func nameTat(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {

	// Constants
	// const name string = "vins"
	// Variables
	// var name1 string = "whos" // or
	// name2 := "vins"
	lengs, upperName := nameTat("vins")
	fmt.Println(plus(2, 2))
	fmt.Println(lengs, upperName)
}
