package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m1 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m)
	fmt.Println(m1, m3)
	for k:= range m {
		fmt.Println(k)
	}
	courseName := m["course"]
	fmt.Println(courseName)

	courseName, ok := m["caurse"]
	fmt.Println(courseName, ok)

	if courseName, ok = m["course"]; ok {
		fmt.Println("course = ", courseName)
	}
	delete(m, "name")
}
