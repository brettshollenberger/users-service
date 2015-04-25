package handlers

import (
	"github.com/brettshollenberger/thrift-example/gen-go/users"
	"github.com/brettshollenberger/users-service/models"
)

type UsersHandler struct {
	usersModel models.UsersModelInterface
}

func NewUsersHandler() *UsersHandler {
	usersModel := models.NewUsersModel()

	return &UsersHandler{
		usersModel,
	}
}

func (this *UsersHandler) Find() (*users.User, error) {
	user, _ := this.usersModel.Find()

	thriftUser := this.usersModel.ToThrift(user)

	return thriftUser, nil
}
