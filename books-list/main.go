package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"gofullstack/books-list/models"
	"log"
)

var books []models.Book

var db *sql.DB

func init()  {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type WithCORS struct {
	r *mux.Router
}

func main()  {
	
}