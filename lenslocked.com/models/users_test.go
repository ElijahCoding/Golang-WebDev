package models

import (
	"fmt"
	"testing"
	"time"
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
	us, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}

	user := User{
		Name: "test",
		Email: "test@test.com",
	}
	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	if user.ID == 0 {
		t.Errorf("Expected ID > 0. Received %d", user.ID)
	}
	if time.Since(user.CreatedAt) > time.Duration(5 * time.Second) {
		t.Errorf("Expected CreatedAt to be recent. Received %s", user.CreatedAt)
	}
	if time.Since(user.UpdatedAt) > time.Duration(5 * time.Second) {
		t.Errorf("Expected UpdatedAt to be recent. Received %s", user.UpdatedAt)
	}
}