package main

import (
	"net/http"
	"os"

	raven "github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

// MailingListRoutes The routes that the core plugins use to determine if they are on the latest version
func MailingListRoutes(r *gin.RouterGroup) {
	r.POST("", mailingListAddUser)
}

// MailingListRequest This is the request made to add a user to the mailing list
type MailingListRequest struct {
	Address string `json:"email"`
}

// @Summary Adds a user to the mailing list
// @Description This adds a user to the mailing list
// @Tags mailing-list
// @Accept  json
// @Produce  json
// @Success 201 {object} main.GenericResponse
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Router /mailing-list [post]
func mailingListAddUser(c *gin.Context) {
	var req MailingListRequest
	var res GenericResponse
	mg := mailgun.NewMailgun("hello.hackerlog.io", os.Getenv("MAILGUN_API_KEY"), "")

	c.BindJSON(&req)

	member := mailgun.Member{
		Address:    req.Address,
		Subscribed: mailgun.Subscribed,
	}

	if err := mg.CreateMember(true, "coming-soon@hello.hackerlog.io", member); err != nil {
		raven.CaptureError(err, map[string]string{
			"email": req.Address,
		})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res.Success = true

	c.JSON(http.StatusCreated, res)
}
