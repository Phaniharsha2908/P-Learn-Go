package main

import "fmt"

type money float64

func (m money) Currency(dollar money) {
	m = dollar
	fmt.Println(m)
}

func main() {
	var dollar money = 70.4
	var m1 money
	m1.Currency(dollar)
}
