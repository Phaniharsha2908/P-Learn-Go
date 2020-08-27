package main

import "fmt"

func main() {

	const (
		m = iota * 2+1
		tue
		wed
		th
		fri
		sat
		sun
	)
	fmt.Println(m, tue, wed, th, fri, sat, sun)

}
