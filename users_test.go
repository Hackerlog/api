package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var uUser = User{
	Email:     "test@test.com",
	FirstName: "Test",
	LastName:  "Dummy",
	Password:  "testing",
	Username:  "test",
}

func TestAuthShouldCreateUser(t *testing.T) {
	u := uUser
	SetupTestDb(&u)

	data, _ := json.Marshal(User{
		Email:     "new@test.com",
		Password:  "password",
		FirstName: "Bob",
		LastName:  "Plunder",
		Username:  "root",
	})
	req, _ := http.NewRequest("POST", "/v1/users", bytes.NewBuffer(data))
	h := SetupTestRouter()
	z := PerformTestRequest(h, req)

	assert.Equal(t, http.StatusCreated, z.Code)
}
