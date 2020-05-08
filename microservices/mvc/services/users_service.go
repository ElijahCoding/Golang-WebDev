package services

import (
	"gofullstack/microservices/mvc/domain"
	"gofullstack/microservices/mvc/utils"
)

type usersService struct {}

var (
	UsersService usersService
)

func (*usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}