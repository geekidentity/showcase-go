package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
	really()
}

/**
 真正闭包
 */
func really()  {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var n int
		n, a = a(i)
		fmt.Println(n)
	}
}
