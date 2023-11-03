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

	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 7

	fmt.Println(len(meuArray))
}
