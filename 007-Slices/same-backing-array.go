package main

import "fmt"

func main() {

	a := []int{34, 7, 9, 10} //abc
	var b []int              // xyz
	b = a[0:2]               // b = abc

	b[1] = 70
	// b[3] = 100 //error
	fmt.Println("a", a)
	fmt.Println("b", b)

}
