package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	password            = "password"
	hashedPassword, err = HashPassword(password)
)

func TestHashPassword(t *testing.T) {
	assert.Nil(t, err)
	assert.NotEqual(t, hashedPassword, password)
}

func TestValidateHashedPassword(t *testing.T) {
	assert.True(t, CheckPassword(hashedPassword, password))
}
