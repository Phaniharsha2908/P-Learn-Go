package calc

import "fmt"

func Final(){
	add() // we don't need to export add as it is being used in same package calc
	sub()
	fmt.Println("Hello")

}
