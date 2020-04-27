package controllers

import (
	"gofullstack/books-list/models"
	"log"
)

var books []models.Book

type Controller struct{}

func logFatal(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}