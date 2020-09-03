package main

import "fmt"

func main() {
	a := []int{23, 65, 87}
	//b := []int{23, 65, 87}
	a = nil

	//fmt.Println(a[0])  // error as we dropped the memory location of a by setting it to nil
	if a == nil { // slices can be compared with nil values only
		fmt.Println("a is nil")
	} else {
		fmt.Println("Not nil")
	}
}
