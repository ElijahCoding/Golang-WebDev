package services

import (
	"gofullstack/microservices/mvc/domain"
	"gofullstack/microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}