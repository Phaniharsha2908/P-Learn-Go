package main

import (
	"fmt"
	"os"
)

func open(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {

		return nil, err

	}
	return file, nil
}

func recoverFileNotFound() {
	r := recover()

	if r != nil {
		fmt.Println("Recovered", r)
	}
}

func main() {

	defer recoverFileNotFound()

	_, err := open("abc.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Fine")

}
