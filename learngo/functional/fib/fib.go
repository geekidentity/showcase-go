package fib

import (
	"fmt"
	"io"
	"strings"
)

/*
Fibonacci
 1, 1, 2, 3, 5, 8, 13, ...
 a, b
    a, b
*/
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 100000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}
