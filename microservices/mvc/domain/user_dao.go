package domain

import (
	"fmt"
	"gofullstack/microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Fede", LastName: "Leon", Email: "myemail@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {}

func (*userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:fmt.Sprintf("user %v was not found.", userId),
		StatusCode: http.StatusNotFound,
		Code: "not found",
	}
}