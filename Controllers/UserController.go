package Controllers

import (
	models "first/Model"
	"first/Repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveUser(c *gin.Context) {

	var input models.User

	input.Prepare()

	input = models.User{Name: input.Name, Email: input.Email, Password: input.Password}
	err := c.BindJSON(&input)

	inputError := input.Validate("")
	if inputError != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"Parameter error": inputError.Error(),
		})
		return
	}

	userCreated, err := Repository.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(), //this error is thrown
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": userCreated,
	})
}

func Login(c *gin.Context) {

	var input models.User
	input = models.User{Name: input.Name, Password: input.Password}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}
	token, err := Repository.AuthenticateUser(input.Name, input.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("User name " + input.Name + " or password is incorrect"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return
}
