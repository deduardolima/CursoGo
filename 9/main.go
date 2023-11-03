package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(2, 3, 5, 6, 8, 7, 4, 2, 3, 5, 123, 235, 421, 1005, 2100))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
