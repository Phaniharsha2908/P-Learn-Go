package main

import "fmt"

func main() {

	a := []int{43, 65, 78}
	//var b []int
	fmt.Printf("%p\n", a)
	a = append(a, 56, 87)

	fmt.Printf("%p\n", a)
	//fmt.Printf("%p\n", b)
	fmt.Println("a=", a)
}
