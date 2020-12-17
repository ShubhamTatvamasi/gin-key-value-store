package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test default key value
func TestGetValueEmpty(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get?key=name", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
}

// Test key after putting value
func TestGetValueWithValue(t *testing.T) {

	newKeyValue := "key=name&value=shubham"

	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()

	// Set a key with value
	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/post", strings.NewReader(newKeyValue))
	req1.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w1, req1)

	// Check the value output
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/get?key=name", nil)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, "shubham", w2.Body.String())
}

// Test post method
func TestPostKeyValue(t *testing.T) {

	newKeyValue := "key=name&value=shubham"

	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/post", strings.NewReader(newKeyValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "name=shubham", w.Body.String())
}

// Test Subscribe Key
func TestSubscribeKey(t *testing.T) {

	newKeyValue := "key=name"

	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/subscribe", strings.NewReader(newKeyValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "name", w.Body.String())
}
