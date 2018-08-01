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
	Password           string     `json:"password" binding:"required"`
	EditorToken        string     `json:"editor_token" gorm:"index"`
	Username           string     `json:"username" gorm:"type:varchar(100);unique_index" binding:"required"`
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
	r.GET("/find/:id", findUser)
	r.GET("/units/:id", getUserWithUnits)
	r.GET("/username", usernameIsAvailable)
	r.GET("/email", emailIsAvailable)
	r.GET("", findUserByEditorToken)
	r.POST("", createUser)
}

// @Summary Gets a user by their ID
// @Description Finds a user given their ID as a path param
// @Tags users
// @Accept  json
// @Produce  json
// @param id path int true "User ID"
// @Success 200 {object} main.User
// @Failure 404 {string} string "Not Found"
// @Router /users/{id} [get]
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

// @Summary Gets a user by their editor token
// @Description Finds a user given their editor token as a path param
// @Tags users
// @Accept  json
// @Produce  json
// @param editor body string true "User's Editor Token"
// @Success 200 {object} main.User
// @Failure 404 {string} string "Not Found"
// @Router /users [get]
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

// @Summary Creates/Registers a user
// @Description Creates a user with the body params that are passed in
// @OperationId createUser
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body main.User true "User object: first_name, last_name, email, password, username"
// @Success 201 {object} main.User
// @Failure 401 {string} string "Bad Request"
// @Router /users [post]
func createUser(c *gin.Context) {
	var user User
	db := GetDb()

	c.ShouldBindJSON(&user)

	if err := db.Create(&user).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusCreated, &user)
	}
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

// AvailableResponse The username response object
type AvailableResponse struct {
	IsAvailable bool `json:"is_available"`
}

// @Summary Checks if a username is available
// @Description Checks if a username is available and responds as such
// @Tags users
// @Accept  json
// @Produce  json
// @param q query string false "Username search using q as key"
// @Success 200 {object} main.AvailableResponse
// @Failure 401 {string} string "Bad Request"
// @Router /users/username [get]
func usernameIsAvailable(c *gin.Context) {
	username := c.Query("q")
	var user User
	var res AvailableResponse

	db := GetDb()

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		res.IsAvailable = true
		c.JSON(http.StatusOK, res)
	} else {
		res.IsAvailable = false
		c.JSON(http.StatusOK, res)
	}
}

// @Summary Checks if an email is available
// @Description Checks if an email is available and responds as such
// @Tags users
// @Accept  json
// @Produce  json
// @param q query string false "email search using q as key"
// @Success 200 {object} main.AvailableResponse
// @Failure 401 {string} string "Bad Request"
// @Router /users/email [get]
func emailIsAvailable(c *gin.Context) {
	email := c.Query("q")
	var user User
	var res AvailableResponse

	db := GetDb()

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		res.IsAvailable = true
		c.JSON(http.StatusOK, res)
	} else {
		res.IsAvailable = false
		c.JSON(http.StatusOK, res)
	}
}
