package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	Email string
}

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
}
