package main

import (
	"fmt"
	"sync"
	"time"
)

var cabs = 2
var wg1 sync.WaitGroup

func main() {
	m := &sync.Mutex{}

	names := []string{"Ravi", "Raj", "Dev", "Vipin", "Ankit"}
	for _, name := range names {
		wg1.Add(1)
		go cab(name, m)
	}

	wg1.Wait()
}

func cab(name string, m *sync.Mutex) {
	m.Lock()
	if cabs >= 1 {
		fmt.Println("Cab is available for ", name)
		time.Sleep(1 * time.Second)
		fmt.Println("Cab Confiremed for ", name)
		fmt.Println("Thanks", name)
		cabs--
	} else {
		fmt.Println("Cab is not available for ", name)
	}
	m.Unlock()
	wg1.Done()
}
