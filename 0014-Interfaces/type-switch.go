package main

import "fmt"

func main() {
	checkType(true)
}

func checkType(i interface{}) {
	switch v := i.(type) {

	case int:
		fmt.Println("Int", v)
	case string:
		fmt.Println("string", v)
	case float64:
		fmt.Println("float", v)
	default:
		fmt.Println("Unknown Type")
	}
}
