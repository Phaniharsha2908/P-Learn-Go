package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "567.05"
	f, _ := strconv.ParseFloat(s, 64)
	fmt.Printf("%T", f)
}
