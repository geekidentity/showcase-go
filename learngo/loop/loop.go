package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
)



func forever()  {
	for {
		fmt.Println("abc")
	}
}

func printFile(filename string)  {
	file, error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func apply(op func(int ,int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println(opName)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += i
	}
	return s
}

func swap1(a, b *int)  {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int)  {
	return b, a
}

func main() {
	printFile("abc.txt")
	//forever()
	fmt.Println(apply(pow, 2, 2))
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 2 ,3))
	fmt.Println(sum(1,2,3,4,5,6,7,8,9))
}