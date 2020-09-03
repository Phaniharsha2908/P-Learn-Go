package main

import "fmt"

func main() {

	a := []int{43, 76, 87, 45}
	b := make([]int, len(a), cap(a))
	copy(b, a)

	fmt.Println(a, b)
	fmt.Printf("%p  %p", a, b)

}
