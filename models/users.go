package models

import (
	"github.com/brettshollenberger/thrift-example/gen-go/users"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
}

type UsersModel struct{}

type UsersModelInterface interface {
	Find() (*User, error)
	ToThrift(*User) *users.User
}

func NewUsersModel() *UsersModel {
	return &UsersModel{}
}

func (this *UsersModel) Find() (*User, error) {
	return &User{
		Id:        1,
		FirstName: "Brett",
		LastName:  "Cassette",
		Email:     "brett.cassette@gmail.com",
	}, nil
}

func (this *UsersModel) ToThrift(user *User) *users.User {
	thriftUser := users.NewUser()
	thriftUser.Id = user.Id
	thriftUser.FirstName = user.FirstName
	thriftUser.LastName = user.LastName
	thriftUser.Email = user.Email

	return thriftUser
}
