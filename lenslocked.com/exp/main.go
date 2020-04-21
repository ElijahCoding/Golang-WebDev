package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gofullstack/lenslocked.com/models"
)

const (
	host = "localhost"
	port = 54320
	user = "postgres"
	password = "root"
	dbname = "lenslocked.com"
)

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID uint
	Amount int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}

	defer us.Close()
	us.DestructiveReset()
	user := models.User{
		Name: "John",
		Email: "john@gmail.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	//user, err := us.ByID(1)
	fmt.Println(user)
}