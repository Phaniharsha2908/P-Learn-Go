package main

import "fmt"

func main() {
	a, b, c := 768, 435, 32

	if a > b && a > c {
		fmt.Println("A is greater")
	} else if b > a && b > c {
		fmt.Println("B is greater")
	} else {
		fmt.Println("C is greater")
	}
}
