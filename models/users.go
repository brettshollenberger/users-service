package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
}

type UsersModel struct {
	db gorm.DB
}

type UsersModelInterface interface {
	Find(id int64) (*User, error)
	Create(user *User) (*User, error)
}

func NewUsersModel(db gorm.DB) *UsersModel {
	return &UsersModel{db}
}

func (this *UsersModel) Find(id int64) (*User, error) {
	user := &User{}

	this.db.First(&user, id)

	return user, nil
}

func (this *UsersModel) Create(user *User) (*User, error) {
	this.db.Create(&user)
	return user, nil
}
