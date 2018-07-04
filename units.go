package main

import (
	"net/http"
	"time"

	"github.com/dericgw/blog-api/users"
	"github.com/gin-gonic/gin"
)

// Unit This represents a file that has been edited, or, a "unit" of work
type Unit struct {
	ID           uint       `json:"id" gorm:"primary_key" example:"1"`
	UserID       uint       `json:"user_id"  gorm:"index" example:"1"`
	EditorType   string     `json:"editor_type"`
	ProjectName  string     `json:"project_name"`
	FileName     string     `json:"file_name"`
	LocWritten   int        `json:"loc_written"`
	LocDeleted   int        `json:"loc_deleted"`
	ComputerType string     `json:"computer_type"`
	Os           string     `json:"os"`
	StartedAt    time.Time  `json:"started_at"`
	StoppedAt    time.Time  `json:"stopped_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"-" gorm:"index"`
}

// UnitRoutes Export the routes for Units
func UnitRoutes(r *gin.RouterGroup) {
	r.POST("", createUnit)
	r.GET("", getUnit)
}

// @Summary Gets units of work for a user
// @Description This gets all of the units of work for a specific user. The user is identified by the
// editor key in the header.
// @Tags units
// @Accept  json
// @Produce  json
// @Param X-Hackerlog-EditorToken header string true "X-Hackerlog-EditorToken"
// @Success 200 {array} main.Unit
// @Router /units [get]
func getUnit(c *gin.Context) {
	var units []Unit
	var user users.User

	db := GetDb()
	eToken := c.GetHeader("X-Hackerlog-EditorToken")

	if err := db.Where("editor_token = ?", eToken).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
	}

	if err := db.Where("user_id = ?", &user.ID).Find(&units).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, &units)
	}
}

// Create Creates a unit of work for a user
func createUnit(c *gin.Context) {
	var unit Unit
	db := GetDb()
	c.BindJSON(&unit)

	db.Create(&unit)
	c.JSON(http.StatusCreated, &unit)
}
