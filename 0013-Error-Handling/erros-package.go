package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var FileNotFound error = errors.New("file not found sum.txt")

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, FileNotFound
	}
	return file, nil
}


func main() {
	_, err := openFile("abc.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fine")
}
