package main

import "fmt"

type fb struct {
	name string
	feed string
}

func (f *fb) new(name string, feed string) {
	f.name = name
	f.feed = feed
}

func (f fb) disp() {
	fmt.Println(f)
}

func main() {
	var u1 fb
	var u2 fb
	u1.new("Diwakar", "abcdddf")
	u2.new("Rahul", "retrtyf")
	u1.disp()
	u2.disp()
}
