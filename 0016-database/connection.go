package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type student struct {
	Name  string
	Class int
	Age   int
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Logger
	if err != nil {
		panic(err)
	}

	s1 := student{
		Name:  "Ajay",
		Class: 9,
		Age:   15,
	}

	db.AutoMigrate(&student{})
	err = db.Create(&s1).Error

	if err != nil {
		log.Fatal(err)
	}

}
