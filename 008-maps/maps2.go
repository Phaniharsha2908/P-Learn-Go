package main

import "fmt"

func main() {

	var student map[int][]string = make(map[int][]string)
	student[1] = []string{"Dev", "Abc"}
	student[2] = []string{"veer", "xyz"}
	names := student[1]
	fmt.Println(names[0])
}
