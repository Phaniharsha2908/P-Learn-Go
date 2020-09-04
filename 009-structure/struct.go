package main

import "fmt"

/*

	Name
	age
	class

student
	var name []string
	var age []int
	var class []int


*/

type student struct {
	name  string //fields of struct
	age   int
	class int
}

func main() {

	var s student
	s.name = "Ajay"
	s.age = 13
	s.class = 10

	fmt.Println(s)
	s.class = 11
	fmt.Printf("%#v",s)
}
