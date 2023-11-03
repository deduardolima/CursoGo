package main

import "fmt"

type ID int

const a = "Hello World"

var (
	b bool = true
	c int  = 10
	d string
	e float64 = 1.5587
	f ID      = 1
)

func main() {

	fmt.Printf("O tipo de E é %T", b)
}
