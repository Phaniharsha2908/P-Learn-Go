package main

import "fmt"

func main() {

	//var i uint8 = 256
	//var a uint8 = 254
	//a+=255
	//var i uint8 = a

	var a int8 = 127
	a += 1
	var i int8 = a

	fmt.Println(i)
}
