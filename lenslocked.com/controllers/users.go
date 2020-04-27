package controllers

import (
	"fmt"
	"gofullstack/lenslocked.com/models"
	"gofullstack/lenslocked.com/rand"
	"gofullstack/lenslocked.com/views"
	"net/http"
)

func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView:     views.NewView("bootstrap", "users/new"),
		LoginView:   views.NewView("bootstrap", "users/login"),
		UserService: us,
	}
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	models.UserService
}

// This is used to render the form where a user can
// create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// This is used to process the signup form when a user
// tries to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := &models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.UserService.Create(user); err != nil {
		panic(err)
	}
	u.signIn(w, user)
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := u.UserService.Authenticate(form.Email, form.Password)
	if user == nil {
		fmt.Println(w, "Invalid login credentials.")
		return
	}
	u.signIn(w, user)
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	user := u.UserService.ByRemember(cookie.Value)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	fmt.Fprintln(w, "User is:", user)
}

func (u *Users) signIn(w http.ResponseWriter, user *models.User) {
	origToken := rand.String(12)
	user.SetRememberToken(origToken)
	if err := u.UserService.Update(user); err != nil {
		panic(err)
	}
	cookie := &http.Cookie{
		Name:     "remember_token",
		Value:    origToken,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}