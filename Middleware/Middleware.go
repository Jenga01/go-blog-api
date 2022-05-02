package Middleware

import (
	"first/Authentication"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func VerifyToken(c *gin.Context) {
	token, ok := GetToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error": "User is unauthorized",
		})
		return
	}
	id, username, err := Authentication.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error": "Token is not valid",
		})
		return
	}
	c.Set("id", id)
	c.Set("name", username)
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.Next()
}

func GetToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}
	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}
	return strings.Trim(arr[1], "\n\t\r"), true
}
