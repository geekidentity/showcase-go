package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa, ss = 2, "kk"
)

func variable() {
	a, b := 3, 4
	var s = "abc"
	fmt.Printf("%d %d %q \n", a, b, s)
	fmt.Println(aa, ss)

}

func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func calTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enums() {
	const (
		cpp = iota
		_
		java
		python
		golang
		js
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang, js)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	fmt.Println("Hello world")
	variable()
	euler()
	triangle()
	consts()
	enums()
}
