package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gofullstack/lenslocked.com/controllers"
	"gofullstack/lenslocked.com/models"
	"net/http"
)

const (
	host = "localhost"
	port = 54320
	user = "postgres"
	password = "root"
	dbname = "lenslocked.com"
)

func main() {
	// Create a DB connection string and then use it to
	// create our model services.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	ug, err := models.NewUserGorm(psqlInfo)
	if err != nil {
		panic(err)
	}
	ug.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(ug)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}