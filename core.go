package main

import (
	"encoding/json"
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

func latestVersion(c *gin.Context) {
	var gBody Release
	cv := c.Query("currentVersion")
	os := c.Query("os")
	arch := c.Query("arch")
	url := githubURL + "/repos/Hackerlog/core/releases/latest"

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

	var download string

	if cv != gBody.Tag {
		for _, i := range gBody.Assets {
			if linkOs, linkArch := extractOsAndArch(i.Download); linkOs == os && linkArch == arch {
				download = i.Download
				break
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"latest":   false,
			"download": download,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"latest":   true,
			"download": nil,
		})
	}
}

// This splits the URL to get the OS and Architecture
func extractOsAndArch(f string) (string, string) {
	n := strings.Split(f, "/")
	m := n[len(n)-1]
	o := strings.Split(m, "_")
	a := strings.Split(o[len(o)-1], ".")
	return o[len(o)-2], a[0]
}
