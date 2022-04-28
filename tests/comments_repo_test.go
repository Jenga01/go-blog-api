package tests

import (
	models "first/Model"
	"first/Repository"
	"fmt"
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
)

func TestCreateComment(t *testing.T) {

	err := refreshCommentsTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedOneUserAndOneArticle()
	if err != nil {
		return
	}
	newComment := models.Comment{
		ID:        1,
		Title:     "Comment title 1",
		Content:   "Content content content",
		ArticleID: 1,
	}
	savedUser, err := Repository.CreateComment(newComment)
	if err != nil {
		t.Errorf("Error saving comment: %v\n", err)
		return
	}
	assert.Equal(t, newComment.ID, savedUser.ID)
	assert.Equal(t, newComment.Title, savedUser.Title)
	assert.Equal(t, newComment.Content, savedUser.Content)
	assert.Equal(t, newComment.ArticleID, savedUser.ArticleID)
}

func TestFindCommentByArticleId(t *testing.T) {
	var comment []models.Comment

	//seededArticle, err := seedOneUserAndOneArticle()
	article, err := seedOneUserAndOneArticle()

	_, err = seedComments()

	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	foundComment, err := Repository.FindCommentsByArticleID(comment, "1")
	if err != nil {
		t.Errorf("error retrieving comments by article id: %v\n", err)
		return
	}

	for i := 0; i < len(*foundComment); i++ {
		fmt.Println("found comment Belong to Article ID: ", (*foundComment)[i].ArticleID)
		assert.Equal(t, (*foundComment)[i].ArticleID, article.ID)
	}
}
