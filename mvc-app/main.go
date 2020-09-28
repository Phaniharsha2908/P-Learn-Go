package main

import (
	"mvc-app/controllers"
	"net/http"
)

func main() {

	http.HandleFunc("/users", controllers.GetUser)
	http.HandleFunc("/hello", controllers.Hello)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
