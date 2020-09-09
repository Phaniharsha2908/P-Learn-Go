package main

import "fmt"

func main() {
	name := "Raj"
	updateName(&name)
	fmt.Println(name)
}

func updateName(s *string) {
	*s = "Rahul"

}
