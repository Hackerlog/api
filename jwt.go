package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// JWT This is the JWT object
type JWT struct {
	Token     string
	ExpiresAt int64
}

// TokenClaims - Claims for a JWT access token.
type TokenClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// CreateJwt Creates a JWT
// Instead of accessing the DB here
// 1. Create the JWT and return it
func CreateJwt(id uint, email string, secret string) (JWT, error) {
	expiresAt := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := TokenClaims{
		ID:    string(id),
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "hackerlog.io",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))

	jwToken := JWT{
		Token:     signedToken,
		ExpiresAt: expiresAt,
	}

	return jwToken, err
}

// ParseJwt Parses a JWT and returns the contents
func ParseJwt(tokenString string, secret string) (bool, jwt.MapClaims) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	m := token.Claims.(jwt.MapClaims)

	if token.Valid {
		return true, m
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Debug("That's not even a token")
			return false, nil
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			log.Debug("Expired")
			return false, nil
		} else {
			// Debug.Println("Couldn't handle this token:", ve)
			log.Debug("Couldn't handle this token:")
			return false, nil
		}
	} else {
		log.Debug("Couldn't handle this token:", err)
		return false, nil
	}
}

// GuardRoute Makes sure the user is authenticated when accessing a route
func GuardRoute(c *gin.Context) {
	var token string
	if header := c.GetHeader("Authorization"); len(header) == 0 {
		c.AbortWithStatus(http.StatusForbidden)
	} else if token = strings.Split(header, " ")[1]; len(token) == 0 {
		c.AbortWithStatus(http.StatusForbidden)
	}
	isValid, _ := ParseJwt(token, os.Getenv("JWT_SECRET"))
	if !isValid {
		c.AbortWithStatus(http.StatusForbidden)
	}
	c.Next()
}
