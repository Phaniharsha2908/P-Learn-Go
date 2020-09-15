package main

import "fmt"

// 1. To satisfy an interface you have to use methods not func
// 2. Create methods with exact same signature as of in interface
// 3. Create all methods in struct as specified in interface. Or your struct should satisfy all of the methods of an interface

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.height * r.width
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r Rect) Hello() {
	fmt.Println("Hello")
}

func main() {

	var s Shape

	r := Rect{
		width:  10,
		height: 10,
	}
	s = r
	fmt.Println("Area=", s.Area())
	fmt.Println("Perm=", s.Perimeter())
	fmt.Printf("%T", s)
}
