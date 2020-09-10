package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	file1, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)

	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	defer file.Close()
	defer file1.Close()

}
