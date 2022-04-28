package tests

import (
	"first/Authentication"
	models "first/Model"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	var user models.User
	token, _ := Authentication.GenerateToken(user)

	assert.NotEqual(t, token, "")
}

func TestValidateToken(t *testing.T) {
	var user models.User
	token, _ := Authentication.GenerateToken(user)
	validated, _, _ := Authentication.ValidateToken(token)

	assert.NotEqual(t, validated, token)
}
