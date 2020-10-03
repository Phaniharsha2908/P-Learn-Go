package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"user-auth/model"
	"user-auth/rand"
)

type jsonData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users struct {
	us *model.UserService
}

func NewUsers(us *model.UserService) *Users {
	return &Users{us: us}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var u1 jsonData
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Unmarshal(b, &u1)

	user := model.User{

		Email:    u1.Email,
		Password: u1.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = u.signIn(w, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Successfully Created the account")

}

func (u *Users) signIn(w http.ResponseWriter, user *model.User) error {

	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = u.us.Update(user)
		if err != nil {
			return err
		}
	}

	cookie := http.Cookie{
		Name:  "remember_token",
		Value: user.Remember,
	}

	http.SetCookie(w, &cookie)

	return nil

}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var UserDetail model.User

	json.Unmarshal(b, &UserDetail)

	user, err := u.us.Authenticate(UserDetail.Email, UserDetail.Password)

	if err != nil {
		switch err {
		case model.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address")
		case model.ErrInvalidPassword:
			fmt.Fprintln(w, "Invalid password")
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	fmt.Fprintln(w, "Logged in, success", user)

}
