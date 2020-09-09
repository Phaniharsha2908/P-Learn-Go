package main

import "fmt"

type user struct {
	name    string
	address string
}

func updateValues(u user) user {
	u.name = "diwakar"
	u.address = "jaipur"
	return u
}

func main() {
	u := user{name: "ravi", address: "abc"}
	u = updateValues(u)
	fmt.Println(u)
}
