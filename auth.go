package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Auth These are the users' JWT's
type Auth struct {
	ID        uint `json:"id" gorm:"primary_key"`
	UserID    uint `json:"user_id"`
	user      User
	Token     string    `json:"token"`
	ExpiresAt int64     `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthRoutes Register the routes
func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/login", checkAuth)
}

func checkAuth(c *gin.Context) {
	var user User
	var rUser User
	db := GetDb()

	c.ShouldBindJSON(&rUser)

	if err := db.Where("email = ?", rUser.Email).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		// check password here
		if authed := CheckPassword(user.Password, rUser.Password); authed != true {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			var userToken Auth
			if findTokenErr := db.Model(&user).Related(&userToken).Error; findTokenErr == nil {
				c.JSON(200, gin.H{
					"token":      &userToken.Token,
					"expires_at": &userToken.ExpiresAt,
				})
			} else {
				// Create the JWT
				jwt, err := CreateJwt(user.ID, user.Email, os.Getenv("JWT_SECRET"))

				// Add the token to the DB
				addJwtToDb(&jwt, user.ID)

				if err != nil {
					c.AbortWithError(http.StatusBadRequest, err)
				} else {
					c.JSON(http.StatusOK, gin.H{
						"token":     &jwt.Token,
						"expiresAt": &jwt.ExpiresAt,
					})
				}
			}
		}
	}
}

func addJwtToDb(j *JWT, userID uint) {
	db := GetDb()
	userToken := Auth{
		UserID:    userID,
		Token:     j.Token,
		ExpiresAt: j.ExpiresAt,
	}

	db.Create(&userToken)
}
