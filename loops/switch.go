package main

import "fmt"

func main() {

	var day string
	fmt.Println("Enter day name")
	fmt.Scanln(&day)
	switch day {
	case "monday":
		fmt.Println("This is monday")
		fallthrough
	case "tuesday":
		fmt.Println("This is tuesday")
		fallthrough
	case "wednesday":
		fmt.Println("This is wed")
	default:
		fmt.Println("Sorry ")
	}

}
