package main

import "fmt"

func add(x int, y int) int {

	sum := x + y
	return sum
	//fmt.Println(x + y)
}
func sub(x, y int) int {
	sum := x + y
	return sum
	fmt.Println("Hello")

	return sum
}
func calc(x, y int) {
	fmt.Println(x * y)
}

func main() {
	a, b := 10, 20

	sum := add(a, b)
	sum1 := sub(b, a)
	calc(sum, sum1)
	//fmt.Println(sum, " from main")
}
