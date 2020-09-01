package main

import "fmt"

func main() {

	a := []int{12, 54, 76, 89, 34, 65, 354, 78}

	b := a[2:5]
	fmt.Println(b)
	fmt.Println("Cap", cap(b), "len", len(b))

}
