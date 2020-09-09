package main

import (
	"fmt"
	"log"
	"strconv"
)

func addInt(a, b string) (int64, error) {

	x, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		return 0, err
	}

	y, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		return 0, err
	}
	fmt.Println(x, y, x+y)
	return x + y, nil

}

func main() {

	result, err := addInt("a", "8")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
