package main

import (
	"fmt"
)

type Expense interface {
	CalculateSalary() int
}

type perm struct {
	empId    int
	basicPay int
	pf       int
}

type contract struct {
	empId    int
	basicPay int
}

func (p perm) CalculateSalary() int {
	return p.basicPay + p.pf
}

func (c contract) CalculateSalary() int {
	return c.basicPay
}

func totalExpense(e ...Expense) {
	expense := 0

	for _, v := range e {
		expense = expense + v.CalculateSalary()
	}
	fmt.Println("Total Expense", expense)
}

func main() {

	e1 := perm{
		empId:    101,
		basicPay: 200,
		pf:       20,
	}

	e2 := perm{
		empId:    102,
		basicPay: 300,
		pf:       30,
	}
	e3 := contract{
		empId:    102,
		basicPay: 300,
	}

	totalExpense(e1, e2, e3)

}
