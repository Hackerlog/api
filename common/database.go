package common

import (
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	// Using this for unit tests
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB The database export
var (
	DB     *gorm.DB
	TempDb = "../test.sqlite"
)

// Init Connects to the database
func Init() *gorm.DB {
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
	if _, fErr := os.Create(TempDb); fErr != nil {
		log.Error("The temp database failed to create", fErr)
	}
	testDb, err := gorm.Open("sqlite3", TempDb)
	if err != nil {
		log.Error(err)
	}
	testDb.DB().SetMaxIdleConns(3)
	testDb.LogMode(true)
	DB = testDb
	return DB
}

// CloseTestDB Delete the database after running testing cases.
func CloseTestDB(testDb *gorm.DB) error {
	testDb.Close()
	err := os.Remove(TempDb)
	return err
}
