package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type person struct {
	Name        string          `json:"first_name"`
	Permissions map[string]bool `json:"perms"`
}

func main() {
	data, err := ioutil.ReadFile("user.json")
	if err != nil {
		log.Println(err)
		return
	}
	var person []person
	if err := json.Unmarshal(data, &person); err != nil {
		log.Println(err)
		return
	}

	fmt.Println(person)

}
