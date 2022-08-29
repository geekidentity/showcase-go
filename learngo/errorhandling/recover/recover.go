package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred: ", err)
		} else {
			panic(r)
		}
	}()
	//a := 1
	//b := 0
	//fmt.Println(a / b)
	//panic(errors.New("this is an error"))
	panic(123)
}

func main() {
	tryRecover()
}
