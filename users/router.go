package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackerlog/api/common"
)

// UserRequest This is a user request that does not contain sensitive info
type UserRequest struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	EditorToken string `json:"editor_token"`
}

// Routes Register the routes
func Routes(r *gin.RouterGroup) {
	r.GET("/:id", Find)
	r.GET("/", FindByEditorToken)
	r.POST("", Create)
}

// Find Finds the first user by their id or by their editor_id
func Find(c *gin.Context) {
	var user User
	db := common.GetDb()
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, &user)
	}
}

// FindByEditorToken Finds the first user by their id or by their editor_id
func FindByEditorToken(c *gin.Context) {
	var user User
	db := common.GetDb()
	editor := c.Query("editor")

	if err := db.Where("editor_token = ?", editor).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, &user)
	}
}

// Create Creates a user and returns the user
func Create(c *gin.Context) {
	var user User
	db := common.GetDb()
	c.BindJSON(&user)

	db.Create(&user)
	c.JSON(http.StatusCreated, &user)
}
