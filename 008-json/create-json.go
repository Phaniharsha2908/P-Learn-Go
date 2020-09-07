package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Perms map[string]bool

type user struct {
	// export your fields to encode them
	FirstName string `json:"first_name"`
	Password  string `json:"-"`
	Perms     `json:"perms,omitempty"`
}

func main() {

	users := []user{
		{
			FirstName: "Diwakar",
			Password:  "abc123",
			Perms:     Perms{"admin": true},
		},
		{
			FirstName: "Rahul",
			Password:  "xyz",
		},

		{
			FirstName: "Ravi",
			Password:  "qwerty",
			Perms:     Perms{"write": false},
		},
	}

	//jsonData, err := json.Marshal(users)
	jsonData, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("user.json", jsonData, 666)
	if err != nil {
		log.Println(err)
		return
	}
}
