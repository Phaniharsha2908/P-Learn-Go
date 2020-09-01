package main

import "fmt"

func main() {

	s := "Hye How"
	// 65 -> A
	// 97 -> a
	for _, v := range s {

		fmt.Printf("%c", v)
	}
}
