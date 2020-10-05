package model

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"user-auth/hash"
)

const userPwPepper = "secret-random-string"
const hmacSecretKey = "secret-hmac-key"

var (
	ErrNotFound        = errors.New("model: resource not found")
	ErrInvalidPassword = errors.New("models:incorrect password provided")
)

type UserService struct {
	db   *gorm.DB
	hmac hash.HMAC
}

func NewUserService(connectionInfo string) (*UserService, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(connectionInfo), &gorm.Config{Logger: newLogger})

	if err != nil {
		return nil, err
	}

	hmac := hash.NewHMAC(hmacSecretKey)

	return &UserService{
		db:   db,
		hmac: hmac,
	}, nil

}

func (us *UserService) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""

	return us.db.Create(user).Error
}

func (us *UserService) Authenticate(email, password string) (*User, error) {

	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPwPepper))

	if err != nil {

		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword

		default:
			return nil, err
		}
	}

	return foundUser, nil

}

func (us *UserService) ByEmail(email string) (*User, error) {

	var user User

	db := us.db.Where("email=?", email)
	err := first(db, &user)

	return &user, err

}

func first(db *gorm.DB, dst interface{}) error {

	err := db.First(dst).Error

	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err

}

func (us *UserService) DestructiveReset() error {

	err := us.db.Migrator().DropTable(&User{})
	if err != nil {
		return err
	}
	return us.AutoMigrate()
}

func (us *UserService) AutoMigrate() error {
	if err := us.db.Migrator().AutoMigrate(&User{}); err != nil {
		return err
	}
	return nil
}

func (us *UserService) Update(user *User) error {

	if user.Remember != "" {
		user.RememberHash = us.hmac.Hash(user.Remember)
	}

	return us.db.Save(user).Error

}

func (us *UserService) ByRememberToken(token string) (*User, error) {

	var user User

	hash := us.hmac.Hash(token)
	fmt.Println("********",hash)
	db := us.db.Where("remember_hash=?", hash)
	err := first(db, &user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
