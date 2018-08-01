package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
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

func TestShouldCreateUser(t *testing.T) {
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

func TestAddProfileImageToUser(t *testing.T) {
	u := uUser
	db := SetupTestDb(&u)
	url := "https://res.cloudinary.com/hhz4dqh1x/image/upload/v1533148995/samples/people/smiling-man.jpg"
	data, _ := json.Marshal(struct {
		ProfileImage string `json:"profile_image"`
	}{
		ProfileImage: url,
	})
	id := strconv.FormatUint(uint64(u.ID), 10)
	req, _ := http.NewRequest("PATCH", "/v1/users/"+id, bytes.NewBuffer(data))
	h := SetupTestRouter()
	z := PerformTestRequest(h, req)

	assert.Equal(t, http.StatusOK, z.Code)

	var assertUser User
	if err := db.Where("id = ?", u.ID).First(&assertUser).Error; err != nil {
		t.Fail()
	}

	assert.Equal(t, url, assertUser.ProfileImage)
}
