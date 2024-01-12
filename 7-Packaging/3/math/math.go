package math

var X string ="Hello World"

type math struct{
	a int
	b int
	C int
}

type Math struct{
	A int
	B int
}
func NewMath(a,b int) math {
	 return math{a:a, b:b}
}

func (m Math) Add() int {
	return m.A + m.B
}


