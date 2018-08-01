package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var aUser = User{
	Email:     "test@test.com",
	FirstName: "Test",
	LastName:  "Dummy",
	Password:  "testing",
	Username:  "test",
}

func TestAuthShouldLogUserIn(t *testing.T) {
	u := aUser
	db := SetupTestDb(&u)
	defer db.Close()

	data, _ := json.Marshal(struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}{
		Password: "testing",
		Email:    "test@test.com",
	})
	req, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(data))
	h := SetupTestRouter()
	z := PerformTestRequest(h, req)

	assert.Equal(t, http.StatusOK, z.Code)
}

func TestAuthShouldCreatePasswordResetToken(t *testing.T) {
	var assertUser User
	u := aUser

	db := SetupTestDb(&u)
	defer db.Close()

	data, _ := json.Marshal(struct {
		Email string `json:"email"`
	}{
		Email: "test@test.com",
	})
	req, _ := http.NewRequest("POST", "/v1/auth/password-reset", bytes.NewBuffer(data))
	h := SetupTestRouter()
	z := PerformTestRequest(h, req)

	assert.Equal(t, http.StatusOK, z.Code)

	if err := db.Where("email = ?", aUser.Email).First(&assertUser).Error; err != nil {
		t.Fail()
	}

	assert.NotNil(t, assertUser.PasswordResetToken)
}
