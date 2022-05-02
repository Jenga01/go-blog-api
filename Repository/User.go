package Repository

import (
	"first/Authentication"
	"first/Config"
	models "first/Model"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (*models.User, error) {

	result := Config.DB.Debug().Create(&user)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &user, nil
}

func AuthenticateUser(name string, password string) (string, error) {
	var user models.User
	if res := Config.DB.Where("name = ?", name).Find(&user); res.Error != nil {
		msg := res.Error
		return "", msg
	}
	err := models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "Incorrect password", err
	}
	return Authentication.GenerateToken(user)
}
