package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db = InitTestDB()

func TestCreateTempDatabase(t *testing.T) {
	_, err := os.Stat(TempDb)
	assert.True(t, !os.IsNotExist(err), "The sqlite temp database does not exist")
	assert.NotNil(t, db)
}

func TestCleansUpTempDB(t *testing.T) {
	err := CloseTestDB(db)
	_, fErr := os.Stat(TempDb)
	assert.True(t, os.IsNotExist(fErr), "The sqlite temp database does not exist")
	assert.Nil(t, err)
}
