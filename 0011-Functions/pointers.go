package main

import "fmt"

func main() {
	var i int = 10
	var x *int = &i // var x =&i
	// x:= &i
	var y *int = x
	fmt.Println(*x) // * dereference operator
	*x = 20
	fmt.Println(x, i)
	*y = 30
	fmt.Println(x, i)

}
