package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTempDatabase(t *testing.T) {
	db := InitTestDB()
	assert.NotEmpty(t, db)
}
