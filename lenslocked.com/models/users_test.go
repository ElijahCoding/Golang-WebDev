package models

import (
	"fmt"
	"testing"
)

func testingUserService() (*UserService, error) {
	const (
		host = "localhost"
		port = 54320
		user = "postgres"
		password = "root"
		dbname = "lenslocked_test"
	)
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := NewUserService(psqlInfo)
	if err != nil {
		return nil, err
	}
	us.db.LogMode(false)
	us.DestructiveReset()
	return us, nil
}

func TestCreateUser(t *testing.T)  {
	
}