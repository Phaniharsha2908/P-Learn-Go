package main

import "fmt"

type user struct {
	name string
	age int
}
type person struct {
	name string
}

func main() {

		var u1 user

		var u2 user

		u1.name = "abc"
		//var p person
		if u1 == u2 { // you can compare same type of struct
			fmt.Println("true")
		} else {
			fmt.Println("False")
		}

		// u1.name == u2.name
		//u1.age == u2.age



}
