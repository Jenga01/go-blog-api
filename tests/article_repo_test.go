package tests

import (
	models2 "first/Model"
	"first/Repository"
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
)

func TestFindAllArticlesWithPagination(t *testing.T) {
	err := refreshUserArticlesCommentsTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table %v\n", err)
	}
	_, _, _, err = seedUsersArticlesComments()
	if err != nil {
		log.Fatalf("Error seeding user and post  table %v\n", err)
	}
	var article models2.Article
	articles, err := Repository.GetAllArticles(&article, &models2.Pagination{Limit: 3})
	if err != nil {
		t.Errorf("this is the error getting the posts: %v\n", err)
		return
	}
	assert.Equal(t, len(*articles), 3)
}

func TestFindArticleById(t *testing.T) {
	var article models2.Article
	err := refreshUserArticlesCommentsTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table: %v\n", err)
	}
	post, err := seedOneUserAndOneArticle()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	foundPost, err := Repository.FindArticleById(&article, "1")
	if err != nil {
		t.Errorf("this is the error getting one article: %v\n", err)
		return
	}
	assert.Equal(t, foundPost.ID, post.ID)
	assert.Equal(t, foundPost.Title, post.Title)
	assert.Equal(t, foundPost.Content, post.Content)
}
