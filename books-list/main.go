package main

import (
	"github.com/gorilla/mux"
)

type Book struct {
	ID int `json:id`
	Title string `json:title`
	Author string `json:author`
	Year string `json:year`
}

var book []Book

func main()  {
	router := mux.NewRouter()
	
}