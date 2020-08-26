package main

import (
	"fmt"
	"os"
)

func main() {
	// 0-> pro-name 1- and so on are arguments
	cmd := os.Args[1:5] // last index = last index -1

	//cmd1 := os.Args[1:] //going to take all of command line arguments



	fmt.Println(cmd)


}
