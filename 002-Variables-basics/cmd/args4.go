package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := os.Args[1:]


	fmt.Println("os.Args", len(os.Args))
	fmt.Println("local var", len(cmd))
}
