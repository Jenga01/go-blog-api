package tests

import (
	models "first/Model"
	"first/Repository"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/pkg/errors"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	newUser := models.User{
		ID:       1,
		Email:    "test@gmail.com",
		Name:     "test",
		Password: "password",
	}
	savedUser, err := Repository.CreateUser(newUser)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Name, savedUser.Name)
}

func TestUserAuthentication(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}

	samples := []struct {
		name     string
		password string
		Message  string
	}{
		{
			name:     user.Name,
			password: "password",
			Message:  "Password match",
		},
	}

	for _, v := range samples {

		token, err := Repository.AuthenticateUser(v.name, v.password)
		if err != nil {
			assert.Equal(t, err, errors.New(v.Message))
		} else {
			assert.NotEqual(t, token, "")
			t.Log(v.Message)
		}
	}
}
