package main

import "fmt"

// 1. To satisfy an interface you have to use methods not func
// 2. Create methods with exact same signature as of in interface
// 3. Create all methods in struct as specified in interface. Or your struct should satisfy all of the methods of an interface

type shape interface {
	Area() float64
	Perimeter() float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) Area() float64 {
	return r.height * r.width
}

func (r rect) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r rect) Hello() {
	fmt.Println("Hello")
}

type square struct {
	width  float64
	height float64
}

func (s square) Area() float64 {
	return s.height * s.width
}

func (s square) Perimeter() float64 {
	return s.height + s.width
}

func main() {

	var s shape // default value is nil

	r := rect{
		width:  10,
		height: 10,
	}
	sq := square{
		width:  15,
		height: 15,
	}
	s = r
	fmt.Println("Area Rect=", s.Area())
	fmt.Println("Perm Rect=", s.Perimeter())
	fmt.Printf("%T\n", s)

	s = sq
	fmt.Println("Area Sq=", s.Area())
	fmt.Println("Perm Sq=", s.Perimeter())
	fmt.Printf("%T", s)
}
