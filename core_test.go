package main

import (
	"net/http/httptest"
	"net/http"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	CoreRoutes(v1.Group("/core"))
	return r
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestIsLatestVersionRoute(t *testing.T) {
	r := setupRouter()

	w := performRequest(r, "GET", "/v1/core/version?currentVersion=v0.04&os=darwin")

	assert.Equal(t, http.StatusOK, w.Code)
}
