package Controllers

import (
	models "first/Model"
	"first/Repository"
	utils "first/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var comment []models.Comment

func SaveComment(c *gin.Context) {

	var input models.Comment

	err := c.BindJSON(&input)
	commentCreated, err := Repository.CreateComment(models.Comment{Title: input.Title, Content: utils.ReplaceBase64ToDecodedImage(input.Content), ArticleID: input.ArticleID})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(), //this error is thrown
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"data": commentCreated,
	})
}

func GetCommentsForArticle(c *gin.Context) {

	commentsList, err := Repository.FindCommentsByArticleID(comment, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i := 0; i < len(*commentsList); i++ {
		c.JSON(http.StatusOK, gin.H{
			"Content": (*commentsList)[i].Content,
			"ID":      (*commentsList)[i].ID,
		})
	}
}
