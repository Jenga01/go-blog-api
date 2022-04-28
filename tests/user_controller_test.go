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

func TestSaveUser(t *testing.T) {

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
			inputJSON:    `{"name":"Vardenis", "email": "test@gmail.com", "password": "password"}`,
			statusCode:   200,
			name:         "Pet",
			email:        "pet@gmail.com",
			errorMessage: "",
		},
	}

	for _, v := range samples {
		r := gin.Default()
		req, err := http.NewRequest("POST", "/register", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		r.POST("/register", Controllers.SaveUser)
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

func TestLogin(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		return
	}

	_, err = seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		name         string
		password     string
		errorMessage string
	}{
		{
			inputJSON:    `{"name": "Vardenis", "password": "password"}`,
			statusCode:   200,
			errorMessage: "",
		},
	}

	for _, v := range samples {

		r := gin.Default()
		req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		r.POST("/login", Controllers.Login)
		r.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 200 {
			assert.NotEqual(t, rr.Body.String(), "")
		}

		if v.statusCode == 422 && v.errorMessage != "" {
			responseMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to json: %v", err)
			}
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}
