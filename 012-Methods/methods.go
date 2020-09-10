package main

import "fmt"

type User struct {
	name  string
	class int
}

func (u User) Disp() {
	u.name = "abc"
	fmt.Println(u)
}

func main() {
	u1 := User{
		name:  "Diwakar",
		class: 9,
	}
	u2 := User{
		name:  "Rahul",
		class: 10,
	}
	u1.Disp()
	u2.Disp()
	fmt.Println("from main", u1)
}
