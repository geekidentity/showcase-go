package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	if contents, error := ioutil.ReadFile(filename); error != nil {
		fmt.Println(error)
	}	else {
		fmt.Printf("%s",contents)
	}
}

func grade(score int) string {
	var result string = ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score: %d", score))
	case score < 60:
		result = "F"
	case score < 100:
		result = "A"
	}
	return result
}