package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

/**
1. defer确保在函数结束时调用
2. 参数在defer语句时计算
3. defer列表为后进先出

何时使用defer
1. open/close
2. lock/unlock
3. PrintHeader/PrintFooter
*/
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}

func deferOrder() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("print too many")
		}
	}
}

func writeFile(filename string) {
	file, error := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if error != nil {
		if pathError, ok := error.(*os.PathError); !ok {
			panic(error)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	if file != nil {
		panic(error)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	deferOrder()
	tryDefer()
	writeFile("fib.txt")
}
