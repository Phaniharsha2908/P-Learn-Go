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
	var persons []person
	if err := json.Unmarshal(data, &persons); err != nil {
		log.Println(err)
		return
	}

	fmt.Println(persons)

	for _, p := range persons {
		fmt.Println(p.Name)

		switch perm := p.Permissions; {
		case perm == nil:
			fmt.Println("has no power")
		case perm["admin"]:
			fmt.Println("is an admin")
		case perm["write"]:
			fmt.Println("has write power")
		}
	}

}
