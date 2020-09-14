package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)
//var err error = fmt.Errorf("any error")
var FileNotFound error = errors.New("file not found abc.txt")

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		customErr := fmt.Errorf("%v %v", err, FileNotFound)
		return nil, customErr
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
