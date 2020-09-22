package main

import "fmt"

func main() {

	c := make(chan int, 2)
	c <- 1
	c <- 2

	a := <-c

	_ = a
	c <- 3

	x, y := <-c, <-c

	fmt.Println(x, y)
}
