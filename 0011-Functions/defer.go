package main

import "fmt"

func main() {
	defer fmt.Println("Bye")
	defer fmt.Println("Bye1")
	fmt.Println("Hello")
	fmt.Println("Hye")
}
