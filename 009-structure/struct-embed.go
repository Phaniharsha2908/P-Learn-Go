package main

import "fmt"

func main() {
	type month int

	type salary struct {
		pf      int
		basePay int
	}

	type fullTimeEmp struct {
		empId        int
		name         string
		WorkingHours int
		salary
	}

	type ContractEmp struct {
		empId    int
		name     string
		Duration month
		salary   //salary salary
	}

	ft := fullTimeEmp{
		empId:        101,
		name:         "Raj",
		WorkingHours: 8,
		salary: salary{
			2000,
			400,
		},
	}

	ct := ContractEmp{
		empId:    102,
		name:     "John",
		Duration: 2,
		salary: salary{
			1000,
			200,
		},
	}

	fmt.Println(ct.empId, ct.pf)
	fmt.Println(ft.empId, ft.pf)

}
