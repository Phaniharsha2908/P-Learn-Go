package main

import "fmt"

type month int

type fullTimeEmp struct {
	empId        int
	name         string
	WorkingHours int
	sal          salary
}
type ContractEmp struct {
	empId    int
	name     string
	Duration month
	sal      salary
}

type salary struct {
	pf      int
	basePay int
}

func main() {

	ft := fullTimeEmp{
		empId:        101,
		name:         "Raj",
		WorkingHours: 8,
		sal: salary{
			2000,
			400,
		},
	}

	ct := ContractEmp{
		empId:    102,
		name:     "John",
		Duration: 2,
		sal: salary{
			1000,
			200,
		},
	}

	fmt.Println(ct.empId, ct.sal.pf)
	fmt.Println(ft.empId, ft.sal.pf)

}
