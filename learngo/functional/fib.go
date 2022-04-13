package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 1, 1, 2, 3, 5, 8, 13, ...
 a, b
    a, b
 */
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
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

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main()  {
	f := fibonacci()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 5
	fmt.Println(f()) // 8
	fmt.Println(f()) // 13
	fmt.Println(f()) // 21
	fmt.Println(f()) // 34
	fmt.Println(f()) // 55
	fmt.Println(f()) // 89
	fmt.Println(f()) // 144
	printFileContents(f)

}