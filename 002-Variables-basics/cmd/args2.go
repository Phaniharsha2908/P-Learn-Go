package main

import (
	"fmt"
	"os"
)

func main() {

	cmd := os.Args[0:2]
	cmd1 := os.Args[2:]

	fmt.Println(cmd, cmd1)

}
