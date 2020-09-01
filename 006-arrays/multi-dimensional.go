package main

import "fmt"

func main() {
	a := [2][3]int{
		{3, 67},      // a[0]
		{54, 87, 78}, // a[1]
	}
	//fmt.Println(a[0])    // complete array at index 0
	//fmt.Println(a[0][1]) // specific value

	for _, v := range a {
		fmt.Println(v)
		for _, val := range v {
			fmt.Println(val)
		}
	}
}
