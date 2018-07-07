package main

import (
	"strings"
	"net/http"
	"encoding/json"
	
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var githubURL = "https://api.github.com"

// CoreRoutes The routes that the core plugins use to determine if they are on the latest version
func CoreRoutes(r *gin.RouterGroup) {
	r.GET("/version", isLatestVersion)
}

type Assets struct {
	Download string `json:"browser_download_url"`
}

type ReleaseResponse struct {
	Tag string `json:"tag_name"`
	Assets []Assets `json:"assets"`
}

func isLatestVersion(c *gin.Context) {
	var gBody ReleaseResponse
	cv := c.Query("currentVersion")
	os := c.Query("os")
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
			if parseFilename(i.Download) == os {
				download = i.Download
				break
			}
		}
		
		c.JSON(http.StatusOK, gin.H{
			"latest": false,
			"download": download,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"latest": true,
			"download": nil,
		})
	}
}

func parseFilename(f string) string {
	n := strings.Split(f, "/")
	m := n[len(n) - 1]
	o := strings.Split(m, "_")
	return o[len(o) - 2]
}
