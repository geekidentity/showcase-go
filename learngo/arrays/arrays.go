package main

import "fmt"

func main() {
	array1()
}

func printArray(arr [5]int)  {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int)  {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func array1()  {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}
	var grid [4][5]int
	fmt.Println(arr1, arr2 ,arr3)
	fmt.Println(grid)
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	printArray(arr1)
	// printArray(arr2)
	printArray(arr3)
	printArray2(&arr1)
	s := make([]int, 16, 32)
	copy(s, arr3[:])
	s = append(s[:3], s[4:]...)
	fmt.Println(s)
	front := s[0]
	s = s[1:]
	tail := s[len(s) - 1]
	s = s[:len(s) - 1]
	fmt.Println(front, tail)
}
