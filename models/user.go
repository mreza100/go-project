package models

import (
	"Waringin/utils/token"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"varchar(300); unique" json:"email" binding:"required"`
	Password string `gorm:"varchar(300)" json:"password" binding:"required"`
	Name     string `gorm:"varchar(300)" json:"name"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(uint(u.Id))

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(*gorm.DB) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil

}

func (u *User) BeforeUpdate(*gorm.DB) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil

}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}
