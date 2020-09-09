package main

import "fmt"

func main() {
	var l [4]int
	updateArray(&l)
	fmt.Println(l)
}

func updateArray(list *[4]int) {
	index := 0
	for i := 10; index < len(list); i += 2 {
		list[index] = i
		index++
	}

}
