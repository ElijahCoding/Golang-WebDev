package main

import (
	"github.com/gorilla/mux"
	"gofullstack/lenslocked.com/controllers"
	"gofullstack/lenslocked.com/views"
	"net/http"
)

var (
	homeView *views.View
	contactView *views.View
)



func main()  {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap","views/contact.gohtml")

	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}

func must(err error)  {
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}