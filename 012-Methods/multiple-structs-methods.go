package main

import "fmt"

type emp struct {
	name string
}

type game struct {
	name string
}

func (e *emp) New(name string) {
	e.name = name
}

func (g *game) New(name string) {
	g.name = name
}

func main() {

	var e1 emp
	e1.New("Amit")
	var g game
	g.New("Mario")

	fmt.Println(e1)
	fmt.Println(g)

}
