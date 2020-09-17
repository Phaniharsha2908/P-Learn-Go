package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go hello(i)
	}
	fmt.Println("Main")
	wg.Wait()
}

func hello(i int) {
	fmt.Println("Hello ", i)
	wg.Done()
}
