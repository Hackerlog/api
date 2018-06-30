package units

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackerlog/api/common"
)

// Routes Export the routes for Units
func Routes(r *gin.RouterGroup) {
	r.POST("", Create)
	r.GET("", Get)
}

// Get Gets all units of work
func Get(c *gin.Context) {
	var units []Unit
	db := common.GetDb()

	if err := db.Find(&units).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, &units)
	}
}

// Create Creates a unit of work for a user
func Create(c *gin.Context) {
	var unit Unit
	db := common.GetDb()
	c.BindJSON(&unit)

	db.Create(&unit)
	c.JSON(http.StatusCreated, &unit)
}
