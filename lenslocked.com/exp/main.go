package main

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"strings"
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
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})

	var u User
	db = db.Where("email = ?", "bla@bla.com").First(&u)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(u)
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("what is your name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("what is your email adddress?")
	email, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	return name, email
}
