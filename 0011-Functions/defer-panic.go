package main

import "fmt"

func main() {
	defer fmt.Println("I want to close a file")
	fmt.Println("Hello")
	panic("There is a problem")
	fmt.Println("Bye")
}
