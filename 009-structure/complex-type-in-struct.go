package main

import "fmt"

func main() {

	type class struct {
		name          []string
		totalStudents int
		class         int
	}

	s1 := class{
		name:          []string{"Dev", "Xyz", "abc"},
		totalStudents: 3,
		class:         5,
	}

	fmt.Println(s1.name)
	for i, v := range s1.name {
		fmt.Println(i, v)
	}

}
