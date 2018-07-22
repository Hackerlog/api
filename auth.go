package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
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
	r.POST("/password-reset", passwordReset)
}

// @Summary Authenticates a user
// @Description Authenticates a user and returns a JWT on successful login
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} main.Auth
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Router /auth/login [post]
func checkAuth(c *gin.Context) {
	var user User
	var rUser User
	db := GetDb()

	c.ShouldBindJSON(&rUser)

	if err := db.Where("email = ?", rUser.Email).First(&user).Error; err != nil {
		log.Debug("No user found")
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		// check password here
		if authed := CheckPassword(user.Password, rUser.Password); authed != true {
			log.Debug("Password does not match")
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			var userToken Auth
			if findTokenErr := db.Model(&user).Related(&userToken).Error; findTokenErr == nil {
				log.Debug("Token found")
				c.JSON(200, gin.H{
					"user":       user,
					"token":      &userToken.Token,
					"expires_at": &userToken.ExpiresAt,
				})
			} else {
				// Create the JWT
				jwt, err := CreateJwt(user.ID, user.Email, os.Getenv("JWT_SECRET"))

				// Add the token to the DB
				addJwtToDb(&jwt, user.ID)

				if err != nil {
					log.Debug("Something happened creating the JWT")
					c.AbortWithError(http.StatusBadRequest, err)
				} else {
					log.Debug("Token created and stored in DB")
					c.JSON(http.StatusOK, gin.H{
						"user":      user,
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

// ResetRequest This is the reset password request
type ResetRequest struct {
	Email string `json:"email"`
}

func passwordReset(c *gin.Context) {
	var req ResetRequest
	var user User
	var res GenericResponse
	res.Error = ""
	res.Success = true

	c.ShouldBindJSON(&req)

	db := GetDb()

	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, res)
	} else {
		token := uuid.New()
		user.PasswordResetToken = token
		db.Save(&user)
		if _, err = sendResetEmail(user.Email, token); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, res)
	}
}

func sendResetEmail(email string, resetKey string) (string, error) {
	template, err := ioutil.ReadFile("emails/reset-password.html")
	if err != nil {
		log.Error("Could not read email template file", err)
	}
	sTemplate := string(template)

	mg := mailgun.NewMailgun("hello.hackerlog.io", os.Getenv("MAILGUN_API_KEY"), "")
	m := mg.NewMessage(
		"Hackerlog Robot <robot@hello.hackerlog.io>",
		"Password Reset",
		"Password Reset",
		email,
	)
	html := strings.Replace(sTemplate, "%URL%", os.Getenv("APP_URL")+"/v1/auth/reset-password/"+resetKey, 1)
	m.SetHtml(html)
	_, id, err := mg.Send(m)
	if err != nil {
		log.Error("Password Reset email did not send", err)
	}
	return id, err
}
