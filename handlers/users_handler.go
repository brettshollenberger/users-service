package handlers

import (
	"github.com/brettshollenberger/thrift-example/gen-go/users"
	"github.com/brettshollenberger/users-service/models"
	"github.com/jinzhu/gorm"
)

type UsersHandler struct {
	usersModel models.UsersModelInterface
}

func NewUsersHandler(db gorm.DB) *UsersHandler {
	usersModel := models.NewUsersModel(db)

	return &UsersHandler{
		usersModel,
	}
}

func (this *UsersHandler) Find(id int64) (*users.User, error) {
	user, _ := this.usersModel.Find(id)

	thriftUser := this.ToThrift(user)

	return thriftUser, nil
}

func (this *UsersHandler) Create(thriftUser *users.User) (*users.User, error) {
	user := this.FromThrift(thriftUser)

	user, _ = this.usersModel.Create(user)

	return thriftUser, nil
}

func (this *UsersHandler) FromThrift(thriftUser *users.User) *models.User {
	user := &models.User{}

	user.FirstName = thriftUser.FirstName
	user.LastName = thriftUser.LastName
	user.Email = thriftUser.Email

	return user
}

func (this *UsersHandler) ToThrift(user *models.User) *users.User {
	thriftUser := users.NewUser()
	thriftUser.Id = &user.Id
	thriftUser.FirstName = user.FirstName
	thriftUser.LastName = user.LastName
	thriftUser.Email = user.Email

	return thriftUser
}
