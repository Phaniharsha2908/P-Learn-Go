package main

import (
	"fmt"
	"log"
	"os"
)

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
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
