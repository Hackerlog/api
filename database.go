package main

import (
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	// Using this for unit tests
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB The database export
var DB *gorm.DB

// DbInit Connects to the database
func DbInit() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error(err)
	}

	if os.Getenv("APP_ENV") == "local" {
		db.LogMode(true)
	}

	DB = db
	return DB
}

// GetDb This allows to pass the database around
func GetDb() *gorm.DB {
	return DB
}

// InitTestDB This function will create a temporarily database for running testing cases
func InitTestDB() *gorm.DB {
	testDb, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Error(err)
	}
	testDb.DB().SetMaxIdleConns(3)

	if os.Getenv("APP_ENV") == "local" {
		testDb.LogMode(true)
	}

	DB = testDb
	return DB
}

// CloseTestDB Delete the database after running testing cases.
func CloseTestDB(testDb *gorm.DB) {
	testDb.Close()
}
