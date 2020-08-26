package main

import "fmt"

func main() {
	// go is a statically typed language
	var i = 30.4 // fix type as float
	// i = "Hello"  float = string // not possible
	i = 30.5
	a := 20

	fmt.Println(i, a)
	fmt.Printf("%T", i)

}
