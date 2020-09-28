
package main

import "fmt"

func main() {

	type student []string

	type class struct {
		name          student
		totalStudents int
		class         int
	}

	s1 := class{
		name:         student{"Dev", "Xyz", "abc"},
		totalStudents: 3,
		class:         5,
	}

	fmt.Println(s1.name)
	for i, v := range s1.name {
		fmt.Println(i, v)
	}

}
