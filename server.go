package main

import (
	"os"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"

	"github.com/hackerlog/api/auth"
	"github.com/hackerlog/api/common"
	"github.com/hackerlog/api/units"
	"github.com/hackerlog/api/users"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var env = os.Getenv("APP_ENV")

func migrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&auth.Auth{})
	db.AutoMigrate(&units.Unit{})
	log.Debug("Migrated DB")
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Debug("No port found in the .env file")
		return ":8000"
	}
	return ":" + port
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Error("No .env file found", err)
	} else {
		log.Debug("Loaded .env file")
	}

	if env == "production" {
		log.Info("Env is in production mode")
		log.SetLevel(log.ErrorLevel)
	} else {
		log.SetLevel(log.DebugLevel)
		log.Debug("Logging everything!!")
	}
}

func main() {
	db := common.Init()
	defer db.Close()

	migrate(db)

	r := gin.Default()

	v1 := r.Group("/v1")

	auth.Routes(v1.Group("/auth"))
	users.Routes(v1.Group("/users"))
	units.Routes(v1.Group("/units"))

	r.Run(getPort())
}
