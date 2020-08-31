package main

import "fmt"

func main() {

	type name [3]string
	// type must be same to compare two arrays
	a := name{"abc", "xyz"}
	b := name{}

	
	if a == b {
		fmt.Println("Equal")
	} else {
		fmt.Println("not")
	}

	fmt.Printf("%#v", a)

}
