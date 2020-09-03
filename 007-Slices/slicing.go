package main

func main() {

	a := []int{345, 54, 67, 87, 34, 6, 765, 5}
	b := a[:4] // starts from 0 location
	c := a[:]
	_, _ = b, c

}
