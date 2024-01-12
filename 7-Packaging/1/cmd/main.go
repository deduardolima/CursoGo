package main

import (
	"fmt"

	"github.com/deduardo/goexpert/7-Packaging/1/math"
)


func main(){
	m := math.Math{A:58,B:45}
	fmt.Println(m.Add())
}