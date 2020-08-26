package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	cmd := os.Args[1:]

	if len(cmd) <= 1 {
		fmt.Println("Sorry Not Enough Values")
		fmt.Println("Please provide your name and age")
		return
	}

	name, ageString := cmd[0], cmd[1]

	age, err := strconv.Atoi(ageString)

	if err != nil {
		fmt.Println(err)
		return
	}

	if age > 16 {
		fmt.Println("Welcome", name, "you can take admission")
	} else {
		fmt.Println("Hye", name, "You must be older than 16 to take admission")
	}

}
