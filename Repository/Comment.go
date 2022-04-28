package Repository

import (
	"first/Config"
	models "first/Model"
)

func CreateComment(comment models.Comment) (*models.Comment, error) {
	result := Config.DB.Debug().Create(&comment)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}

	return &comment, nil
}

func FindCommentsByArticleID(comments []models.Comment, ArticleId string) (*[]models.Comment, error) {

	result := Config.DB.Model(&models.Comment{}).Where("article_id = ?", ArticleId).Find(&comments)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &comments, nil
}
