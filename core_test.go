package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func vReq(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	req.Header.Add(xHeader, TestEditorToken)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestIsLatestVersionRoute(t *testing.T) {
	var user User
	user.Email = "test@test.com"
	user.EditorToken = TestEditorToken
	user.FirstName = "Test"
	user.LastName = "Dummy"
	user.Password = "password"

	SetupTestDb(&user)

	r := SetupTestRouter()
	w := vReq(r, "GET", "/v1/core/version?currentVersion=v0.04&os=darwin")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestParsesURLAndReturnsOSAndArch(t *testing.T) {
	tt := []struct {
		os   string
		arch string
	}{
		{os: "windows", arch: "amd64"},
		{os: "windows", arch: "386"},
		{os: "darwin", arch: "amd64"},
		{os: "darwin", arch: "386"},
		{os: "linux", arch: "amd64"},
		{os: "linux", arch: "386"},
	}
	for _, tc := range tt {
		os, arch := extractOsAndArch(fmt.Sprintf("https://github.com/octocat/Hello-World/releases/download/v1.0.0/core_0.4_%s_%s.tar.gz", tc.os, tc.arch))
		assert.Equal(t, os, tc.os)
		assert.Equal(t, arch, tc.arch)
	}
}
