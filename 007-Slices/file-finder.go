package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("Provide Directory Name")
		return
	}

	abc, err := ioutil.ReadDir(args[0])

	if err != nil {
		log.Println(err)
		return
	}
	var names []byte
	for _, f := range abc{
		if f.Size() == 0 {
			names = append(names, f.Name()...)
			names = append(names, '\n')
		}
	}

	err = ioutil.WriteFile("out.txt", names, 666)

	if err != nil {
		log.Println(err)
		return
	}
}
