package main

import "fmt"

func main() {

	var n []int
	n = make([]int, 4)
	UpdateSlice(n)
	fmt.Println(n)
}

func UpdateSlice(list []int) {
	for i := 0; i < len(list); i++ {
		list[i] = i
	}
}
