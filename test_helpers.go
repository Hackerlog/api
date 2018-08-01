package main

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// TestEditorToken Used for testing
const TestEditorToken = "test"

// SetupTestRouter Router for testing
func SetupTestRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")

	AuthRoutes(v1.Group("/auth"))
	UserRoutes(v1.Group("/users"))
	UnitRoutes(v1.Group("/units"))
	CoreRoutes(v1.Group("/core"))
	MailingListRoutes(v1.Group("/mailing-list"))

	return r
}

// PerformTestRequest Makes test requests
func PerformTestRequest(h http.Handler, r *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w
}

// SetupTestDb DB used for testing
func SetupTestDb(data interface{}) *gorm.DB {
	db := InitTestDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Auth{})
	db.Create(data)
	return db
}
