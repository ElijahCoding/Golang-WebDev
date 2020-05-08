package services

import (
	"gofullstack/microservices/mvc/domain"
	"gofullstack/microservices/mvc/utils"
	"net/http"
)

type itemService struct {
	
}

func (*itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}