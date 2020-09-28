package controllers

import (
	"encoding/json"
	"fmt"
	"mvc-app/model"
	"mvc-app/util"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdString := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseUint(userIdString, 10, 64)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		appErr := util.ApplicationError{
			Message:    "User Id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(appErr)
		w.WriteHeader(appErr.StatusCode)
		w.Write(jsonValue)
		return
	}
	user, appErr := model.UserService.GetUser(userId)

	if appErr != nil {
		jsonValue, _ := json.Marshal(appErr)
		w.WriteHeader(appErr.StatusCode)
		w.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonValue)
	return

}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hye from Hello")
}
