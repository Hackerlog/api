package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hackerlog/api/common"
	"github.com/hackerlog/api/users"
)

// Routes Register the routes
func Routes(r *gin.RouterGroup) {
	r.POST("/login", Check)
}

// Check Is the user registered and are the credentials valid
func Check(c *gin.Context) {
	var user users.User
	var rUser users.User
	db := common.GetDb()

	c.ShouldBindJSON(&rUser)

	if err := db.Where("email = ?", rUser.Email).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		// check password here
		if authed := common.CheckPassword(user.Password, rUser.Password); authed != true {
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
				jwt, err := common.CreateJwt(user.ID, user.Email, os.Getenv("JWT_SECRET"))

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

func addJwtToDb(j *common.JWT, userID uint) {
	db := common.GetDb()
	userToken := Auth{
		UserID:    userID,
		Token:     j.Token,
		ExpiresAt: j.ExpiresAt,
	}

	db.Create(&userToken)
}
