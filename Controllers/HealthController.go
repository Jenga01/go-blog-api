package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {

	c.JSON(http.StatusAccepted, gin.H{
		"data": "Health is ok",
	})

}
