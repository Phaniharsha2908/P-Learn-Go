package main

import "strconv"

func main() {

	var a int
	s := "10"
	a, err := strconv.Atoi(s)
	_, _, abc := a, err, 1
	_ = abc
	a, err = strconv.Atoi(s)
}
