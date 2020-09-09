package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{name: "Amit"}
	updateStruct(&p)
	fmt.Println(p)

	//var p1 *person = &p
	//p1 := &p
	//p1.name = "Raj"

}

func updateStruct(p1 *person) {
	p1.name = "raj"

}
