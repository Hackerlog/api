package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// TODO: Need to add a token to this call at some point
var githubURL = "https://api.github.com"

// CoreRoutes The routes that the core plugins use to determine if they are on the latest version
func CoreRoutes(r *gin.RouterGroup) {
	r.GET("/version", latestVersion)
}

// Assets The download links within a response from Github
type Assets struct {
	Download string `json:"browser_download_url"`
}

// Release The release response from Github
type Release struct {
	Tag    string   `json:"tag_name"`
	Assets []Assets `json:"assets"`
}

// VersionResponse The response returned for the version endpoint
type VersionResponse struct {
	Latest   bool   `json:"latest" example:"false"`
	Download string `json:"download" example:"https://github.com/Hackerlog/core/releases/download/v0.5/core_0.5_windows_amd64.zip"`
}

// @Summary Returns a link of the latest version of the Core app
// @Description This endpoint takes a few parameters and with those parameters, it looks to see if
// the client has the most recent version of the Core app. If it does, it sends a boolean saying so.
// If it does not, then it sends back a boolean saying so, but also a URL to the latest download
// specific to the OS and architecture submitted.
// @Tags core
// @Accept  json
// @Produce  json
// @Param X-Hackerlog-EditorToken header string true "X-Hackerlog-EditorToken"
// @Success 200 {object} main.VersionResponse
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Router /core/version [get]
func latestVersion(c *gin.Context) {
	var gBody Release
	var response VersionResponse
	cv := c.Query("currentVersion")
	os := c.Query("os")
	arch := c.Query("arch")
	url := githubURL + "/repos/Hackerlog/core/releases/latest"

	// Do a little validation here to make sure we have a signed up user
	var user User

	db := GetDb()
	eToken := c.GetHeader(xHeader)

	if err := db.Where("editor_token = ?", eToken).First(&user).Error; err != nil {
		raven.CaptureError(err, map[string]string{
			"editor_token": eToken,
		})
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	resp, err := http.Get(url)

	if err != nil {
		log.Error("The request to get the latest release failed", err)
		raven.CaptureError(err, nil)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&gBody); err != nil {
		log.Error("Failed to parse response from latest release", err)
		raven.CaptureError(err, nil)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if cv != gBody.Tag {
		for _, i := range gBody.Assets {
			if linkOs, linkArch := extractOsAndArch(i.Download); linkOs == os && linkArch == arch {
				fmt.Println("Found it")
				response.Download = i.Download
				break
			}
		}

		response.Latest = false
	} else {
		response.Latest = true
		response.Download = ""
	}

	c.JSON(http.StatusOK, response)
}

// This splits the URL to get the OS and Architecture
func extractOsAndArch(f string) (string, string) {
	n := strings.Split(f, "/")
	m := n[len(n)-1]
	o := strings.Split(m, "_")
	a := strings.Split(o[len(o)-1], ".")
	return o[len(o)-2], a[0]
}
