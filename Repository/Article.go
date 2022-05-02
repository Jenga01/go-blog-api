package Repository

import (
	"first/Config"
	models2 "first/Model"
)

func GetAllArticles(article *models2.Article, pagination *models2.Pagination) (*[]models2.Article, error) {

	var articles []models2.Article
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := Config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuilder.Model(&models2.Article{}).Where(article).Find(&articles)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &articles, nil
}

func FindArticleById(article *models2.Article, pid string) (*models2.Article, error) {

	var err error
	err = Config.DB.Model(&models2.Article{}).Where("id = ?", pid).First(&article).Error
	if err != nil {
		return &models2.Article{}, err
	}
	return article, nil
}
