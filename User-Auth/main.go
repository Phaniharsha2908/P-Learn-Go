package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"user-auth/controllers"
	"user-auth/model"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	dsn := os.Getenv("NAME") + ":" + os.Getenv("PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	us, err := model.NewUserService(dsn)
	if err != nil {
		panic(err)
	}
	//us.DestructiveReset()
	//us.AutoMigrate()

	userC:= controllers.NewUsers(us)

	r := mux.NewRouter()

	r.HandleFunc("/signup", userC.Create).Methods("POST")
	r.HandleFunc("/login",userC.Login)
	err = http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
