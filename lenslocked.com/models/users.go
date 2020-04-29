package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gofullstack/lenslocked.com/hash"
)
// We will get these out of our code later.
// var userPwPepper = "secret-random-string"
var hmac = hash.NewHMAC("secret-hmac-key")
var bcrypt = hash.NewBCrypt("secret-random-string")

type User struct {
	gorm.Model
	Name          string
	Email         string `gorm:"not null;unique_index"`
	Password      string `gorm:"-"`
	PasswordHash  string `gorm:"not null"`
	RememberToken string
}

func (u *User) SetRememberToken(token string) {
	u.RememberToken = hmac.String(token)
}

type UserService interface {
	ByID(id uint) *User
	ByEmail(email string) *User
	ByRemember(token string) *User
	Authenticate(email, password string) *User
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

func NewUserGorm(connectionInfo string) (*UserGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	return &UserGorm{db}, nil
}

type UserGorm struct {
	*gorm.DB
}

func (ug *UserGorm) ByID(id uint) *User {
	return ug.byQuery(ug.DB.Where("id = ?", id))
}

func (ug *UserGorm) ByEmail(email string) *User {
	return ug.byQuery(ug.DB.Where("email = ?", email))
}

func (ug *UserGorm) ByRemember(hashedToken string) *User {
	return ug.byQuery(
		ug.DB.Where("remember_token = ?", hmac.String(hashedToken)))
}

func (ug *UserGorm) byQuery(query *gorm.DB) *User {
	ret := &User{}
	err := query.First(ret).Error
	switch err {
	case nil:
		return ret
	case gorm.ErrRecordNotFound:
		return nil
	default:
		panic(err)
	}
}

func (ug *UserGorm) Authenticate(email, password string) *User {
	foundUser := ug.ByEmail(email)
	if foundUser == nil {
		// No user found with that email address
		return nil
	}

	if !bcrypt.Equal(foundUser.PasswordHash, password) {
		return nil
	}
	return foundUser
}

func (ug *UserGorm) Create(user *User) error {
	user.PasswordHash = bcrypt.String(user.Password)
	user.Password = ""
	return ug.DB.Create(user).Error
}

func (ug *UserGorm) Update(user *User) error {
	return ug.DB.Save(user).Error
}

func (ug *UserGorm) Delete(id uint) error {
	user := &User{Model: gorm.Model{ID: id}}
	return ug.DB.Delete(user).Error
}

func (ug *UserGorm) DestructiveReset() {
	ug.DropTableIfExists(&User{})
	ug.AutoMigrate()
}

func (ug *UserGorm) AutoMigrate() {
	ug.DB.AutoMigrate(&User{})
}