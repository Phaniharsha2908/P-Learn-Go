package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	c := make(chan int)

	wg.Add(4)

	go add(2, 2, c)
	go sub(4, 2, c)
	go multiply(3, 3, c)
	go calc(c)

	wg.Wait()
}

func add(a int, b int, c chan int) {
	sum := a + b
	c <- sum
	wg.Done()
}

func sub(a int, b int, c chan int) {
	sum := a - b
	c <- sum
	wg.Done()
}

func multiply(a int, b int, c chan int) {
	sum := a * b
	c <- sum
	wg.Done()
}

func calc(c chan int) {

	a, b, d := <-c, <-c, <-c
	fmt.Println(a + b + d)
	wg.Done()
}
