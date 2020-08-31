package main

import "fmt"

func main() {

	var i int = 10
	{
		i = i + 20
		fmt.Println("Inner", i)
	}
	fmt.Println(i)
}
