package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
)

// User This is the user model that will hold all of the users
type User struct {
	ID                 uint       `json:"id" gorm:"primary_key"`
	Email              string     `json:"email" gorm:"type:varchar(100);unique_index" binding:"required"`
	FirstName          string     `json:"first_name" binding:"required"`
	LastName           string     `json:"last_name" binding:"required"`
	Password           string     `json:"-" binding:"required"`
	EditorToken        string     `json:"editor_token" gorm:"index"`
	PasswordResetToken string     `json:"-"`
	Units              []Unit     `json:"units"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"-" gorm:"index"`
}

// BeforeCreate We want to hash the users password
func (u *User) BeforeCreate() (err error) {
	hashedPassword, err := HashPassword(u.Password)

	if err != nil {
		err = errors.New("Hashing user password failed")
	}

	u.Password = hashedPassword

	if u.EditorToken == "" {
		u.EditorToken = uuid.New()
	}

	return nil
}

// UserRoutes Register the routes
func UserRoutes(r *gin.RouterGroup) {
	r.GET("/:id", findUser)
	r.GET("/", findUserByEditorToken)
	r.GET("/:id/units", getUserWithUnits)
	r.POST("", createUser)
}

func findUser(c *gin.Context) {
	var user User
	db := GetDb()
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, &user)
	}
}

func findUserByEditorToken(c *gin.Context) {
	var user User
	db := GetDb()
	editor := c.Query("editor")

	if err := db.Where("editor_token = ?", editor).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, &user)
	}
}

func createUser(c *gin.Context) {
	var user User
	db := GetDb()
	c.BindJSON(&user)
	c.BindJSON(&user)

	db.Create(&user)
	c.JSON(http.StatusCreated, &user)
}

func getUserWithUnits(c *gin.Context) {
	var user User

	db := GetDb()
	id := c.Param("id")

	if err := db.Preload("Units").Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, &user)
	}
}
