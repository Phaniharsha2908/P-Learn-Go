package main

import "fmt"

func main() {

	show(23, "abc", 98, 89, 34)

}

func show(i int, s1 string, s ...int) {
	fmt.Printf("%#v\n", i)
	fmt.Printf("%#v\n", s1)
	fmt.Printf("%#v\n", s)
}

func disp(i []int) {
	fmt.Printf("%#v", i)
	
}
