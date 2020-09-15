package main

import "fmt"

func main() {
	var i, v interface{} = "Hello", 34

	fmt.Printf("%#v %#v\n", i, v)
	fmt.Printf("%T %T\n", i, v)
	disp(4,"Hye",true)

}
func disp(i ...interface{})  {
	fmt.Println(i)
}
