package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique"`
}

var db *gorm.DB

func main() {

	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // Slow SQL threshold
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      true,        // Disable color
	//	},
	//)

	dsn := "root:root@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Logger: newLogger

	if err != nil {
		panic(err)
	}
	//CreateTable()
	//InsertData()
	//SearchData()
	//SearchAll()
	//SearchWhere()
	//Update()
	//UpdateWhere()
	Delete()
}

func CreateTable() {
	db.Migrator().DropTable(&user{})
	db.Migrator().AutoMigrate(&user{})
}

func InsertData() {
	users := []user{
		{
			Name:  "raj",
			Email: "raj@gmail.com",
		},
		{
			Name:  "dev",
			Email: "dev@gmail.com",
		},
		{

			Name:  "ravi",
			Email: "ravi@gmail.com",
		},
		{

			Name:  "ravi",
			Email: "ravi1234@gmail.com",
		},
	}

	err := db.Create(&users).Error

	if err != nil {
		fmt.Println(err)
		return
	}

}

func SearchData() {

	var user user
	err := db.First(&user).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)

}

func SearchAll() {
	var users []user

	err := db.Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(users)
}

func SearchWhere() {

	var user1 user
	err := db.Where("name = ?", "ravi").First(&user1).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user1)

	var users []user

	db.Where("name=?", "ravi").Find(&users)
	fmt.Println(users)
}

func Update() {
	var user user
	db.First(&user)

	user.Name ="abc"
	db.Save(&user)
}

func UpdateWhere() {
	db.Model(&user{}).Where("email=?","raj@gmail.com").Update("name","raj")
}

func Delete() {

	var user user
	db.First(&user)

	db.Unscoped().Delete(&user)

}
