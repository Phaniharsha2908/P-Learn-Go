package main

import "fmt"

func main() {

	type student struct {
		name  string //fields of struct
		age   int
		class int
	}

	students := []student{
		{
			name: "Dev",
			age:  12,
		},
		{
			name: "xyz",
			age:  13,
		},
	}
	//fmt.Println(students[0].name)

	for i, s := range students {
		fmt.Println(i,s.name)
	}

}
