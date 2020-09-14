package main

import (
	"errors"
	"fmt"
	"os"
)

var FileNotFound error = errors.New("file not found abc.txt")

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		customErr := fmt.Errorf("%v %w", err, FileNotFound) // you can wrap one value at a time
		return nil, customErr
	}
	return file, nil
}

func main() {
	_, err := openFile("abc.txt")

	if errors.Is(err, FileNotFound) { //==
		fmt.Println("Matches")
	} else {
		fmt.Println("Not")
	}

}
