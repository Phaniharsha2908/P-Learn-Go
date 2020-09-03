package main

import "fmt"

func main() {

	var i []int
	i = make([]int, 5)
	fmt.Println(i)
	fmt.Printf("%p\n", i)

	i = append(i, 54, 76)
	fmt.Printf("%p\n", i)

	i = append(i, 45, 89)
	fmt.Printf("%p\n", i)

	i = append(i, 45, 89)
	fmt.Printf("%p\n", i)

	fmt.Println(i)
}
