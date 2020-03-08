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

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sum(1,2,3,4,5,6)
	fmt.Println(result)
}
