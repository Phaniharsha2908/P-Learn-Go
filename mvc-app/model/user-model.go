package model

import (
	"fmt"
	"mvc-app/util"
	"net/http"
)

var person = map[uint64]*User{
	123: &User{
		FName: "Diwakar",
		LName: "Singh",
		Email: "diwakar@gmail.com",
	},
	124: &User{
		FName: "Ravi",
		LName: "Kumar",
		Email: "ravi@gmail.com",
	},
}

type userService struct{}

var UserService userService


func (us *userService) GetUser(userId uint64) (*User, *util.ApplicationError) {
	u := person[userId]

	if u != nil {
		return u, nil
	}

	return nil, &util.ApplicationError{
		Message:    fmt.Sprintf("User Not Found with user id %v", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
