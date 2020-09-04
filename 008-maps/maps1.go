package main

import "fmt"

func main() {

	class := map[int][]string{
		10: {"abc", "xyz", "qwer"},
	}

	fmt.Println(class[10])
}
