package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	secret                 = "secret"
	id                uint = 1
	email                  = "test@test.com"
	token, tokenError      = CreateJwt(id, email, secret)
)

func TestCreateValidJWT(t *testing.T) {
	assert.Nil(t, tokenError)
}

func TestReturnsTrueIfValidTokenIsPassed(t *testing.T) {
	isValid, _ := ParseJwt(token.Token, secret)
	assert.True(t, isValid)
}

func TestReturnsNilIfSignatureIsInvalid(t *testing.T) {
	isValid, claims := ParseJwt(token.Token, "wrongsignature")
	assert.False(t, isValid)
	assert.Nil(t, claims)
}
