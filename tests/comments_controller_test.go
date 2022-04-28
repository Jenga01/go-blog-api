package tests

import (
	"bytes"
	"encoding/json"
	"first/Controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveComment(t *testing.T) {

	err := refreshUserArticlesCommentsTable()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		name         string
		email        string
		errorMessage string
	}{
		{
			inputJSON:    `{"title":"title 1", "content": "Content content content", "article_id": 1}`,
			statusCode:   200,
			errorMessage: "",
		},
	}

	for _, v := range samples {
		r := gin.Default()
		req, err := http.NewRequest("POST", "/articles/comment", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		r.POST("/articles/comment", Controllers.SaveComment)
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 200 {
			assert.NotEqual(t, rr.Body.String(), "")
		}
		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestGetCommentsByArticleId(t *testing.T) {
	err := refreshUserArticlesCommentsTable()
	if err != nil {
		log.Fatal(err)
	}
	comments, err := seedComments()
	article, _ := seedOneArticle()

	gin.SetMode(gin.TestMode)
	handler := Controllers.GetCommentsForArticle
	router := gin.Default()
	router.GET("/articles/comments/:id", handler)

	req, err := http.NewRequest("GET", "/articles/comments/1", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	for _, v := range comments {
		if http.StatusOK == resp.Code {
			assert.Equal(t, v.ArticleID, article.ID)
		}
	}
}
