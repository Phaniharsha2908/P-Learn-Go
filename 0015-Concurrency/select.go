package main

import "fmt"

func main() {
	c := make(chan string)
	c1 := make(chan string)

	go func() {
		c <- "one"
	}()

	go func() {
		c1 <- "two"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case a := <-c:
			fmt.Println("received ", a)
		case b := <-c1:
			fmt.Println("received ", b)
		}
	}

}
