package tests

import (
	"first/Controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindAllArticles(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	err := refreshUserArticlesCommentsTable()
	if err != nil {
		log.Fatal(err)
	}

	_, articles, _, err := seedUsersArticlesComments()
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/articles/all", Controllers.GetArticles)
	req, err := http.NewRequest(http.MethodGet, "/articles/all", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, len(articles), 4)
	}
	assert.NotEqual(t, len(articles), 2)
}

func TestGetArticleById(t *testing.T) {
	article, _ := seedOneArticle()
	_, _, _, _ = seedUsersArticlesComments()
	articleSamples := []struct {
		id           uint64
		statusCode   int
		title        string
		content      string
		errorMessage string
	}{
		{
			id:         article.ID,
			statusCode: 200,
			title:      article.Title,
			content:    article.Content,
		},
		{
			statusCode: 400,
		},
	}
	for _, v := range articleSamples {
		gin.SetMode(gin.TestMode)
		handler := Controllers.GetArticleById
		router := gin.Default()
		router.GET("/articles/article/:id", handler)

		req, err := http.NewRequest("GET", "/articles/article/1", nil)
		if err != nil {
			fmt.Println(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		if v.statusCode == resp.Code {
			assert.Equal(t, article.ID, v.id)
		}

	}
}
