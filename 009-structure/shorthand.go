package main

import "fmt"

func main() {

	type student struct {
		name  string //fields of struct
		age   int
		class int
	}

	s := student{
		name:  "Dev",
		age:   14,
		class: 9,
	}

	var name string
	name = s.name
	fmt.Println(name)

}
