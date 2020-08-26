package calc

import "fmt"

var X, y int
var a = 10

func add() { // make first letter as uppercase to export it

	X, y = 100, 200
	//sum:=x+y
	fmt.Println(X + y)

}
func sub() {
	X, y = 100, 30
	//sum:=x-y
	fmt.Println(X - y)
}
