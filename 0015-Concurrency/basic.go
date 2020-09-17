package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	go disp() // writing go in front of function call creates a goroutine

	fmt.Println("Hello")

	time.Sleep(5 * time.Millisecond)
}

func disp() {
	fmt.Println("Hello from disp")
}
