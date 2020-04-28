package controllers

import (
	"database/sql"
	"gofullstack/books-list/models"
	"log"
	"net/http"
)

var books []models.Book

type Controller struct{}

func logFatal(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}