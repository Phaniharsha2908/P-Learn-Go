package main

import "fmt"

type money int64 //

type rupee money
type paisa rupee

func main() {
	var r = 1
	var dollar money = 100
	var rupee money = money(r)

	fmt.Println(dollar, rupee)
}
