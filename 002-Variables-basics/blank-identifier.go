package main

import "os"

func main() {

	a := "Hello"
	_ = a

	_, err := os.Open("abc.txt")
	_ = err
}
