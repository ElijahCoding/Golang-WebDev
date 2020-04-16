package controllers

import (
	"fmt"
	"gofullstack/lenslocked.com/views"
	"net/http"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Fprint(w, r.PostForm["email"])
	fmt.Fprint(w, r.PostFormValue("email"))
	fmt.Fprint(w, r.PostForm["password"])
	fmt.Fprint(w, r.PostFormValue("password"))

}