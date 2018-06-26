package users

import (
	"net/http"

	"github.com/dericgw/blog-api/common"
	"github.com/gin-gonic/gin"
)

// Routes Register the routes
func Routes(r *gin.RouterGroup) {
	r.POST("", Create)
}

// Create Creates a user and returns the user
func Create(c *gin.Context) {
	var user User
	db := common.GetDb()
	c.BindJSON(&user)

	db.Create(&user)
	c.JSON(http.StatusCreated, &user)
}
