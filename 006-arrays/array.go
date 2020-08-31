package main

import "fmt"

// it has a fixed size
// all elm must be of same type
func main() {

	var num [5]int
	num[0] = 10
	num[4] = 20
	fmt.Printf("%#v", num)

}
