package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var authors []Author = []Author{
	Author{
		Id:        "author-1",
		Firstname: "Nic",
		Lastname:  "Raboy",
		Username:  "nraboy",
		Password:  "pass",
	},
	Author{
		Id:        "author-2",
		Firstname: "Maria",
		Lastname:  "Raboy",
		Username:  "mraboy",
		Password:  "abc123",
	},
}

var articles []Article = []Article{
	Article{
		Id:      "article-1",
		Author:  "author-1",
		Title:   "Blog Post 1",
		Content: "This is an example blog article",
	},
}

func RootEndpoint(response http.ResponseWriter, request *http.Request)  {
	response.Header().Add("content-type", "application/json")
	response.Write([]byte(`{ "message": "testing" }`))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	router.HandleFunc("/authors", AuthorRetrieveAllEnpoint).Methods("GET")
	router.HandleFunc("/author/{id}", AuthorRetrieveEndpoint).Methods("GET")
	router.HandleFunc("/author/{id}", AuthorDeleteEndpoint).Methods("DELETE")
	router.HandleFunc("/author/{id}", AuthorUpdateEndpoint).Methods("PATCH")

	router.HandleFunc("/articles", ArticleRetrieveAllEndpoint).Methods("GET")
	router.HandleFunc("/article/{id}", ArticleRetrieveEndpoint).Methods("GET")
	router.HandleFunc("/article/{id}", ArticleDeleteEndpoint).Methods("DELETE")
	router.HandleFunc("/article/{id}", ArticleUpdateEndpoint).Methods("PUT")
	router.HandleFunc("/article", ArticleCreateEndpoint).Methods("POST")

	http.ListenAndServe(":8080", router)
}