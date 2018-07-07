package main

import (
	"os"

	_ "github.com/hackerlog/api/docs"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	
	log "github.com/sirupsen/logrus"
	raven "github.com/getsentry/raven-go"
	"github.com/evalphobia/logrus_sentry"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var env = os.Getenv("APP_ENV")

func migrate(db *gorm.DB) {
	db.AutoMigrate(&Unit{})
	db.AutoMigrate(&Auth{})
	db.AutoMigrate(&User{})
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

	client, err := raven.New(os.Getenv("SENTRY_DSN"))

	if err != nil {
		log.Fatal(err)
	}

	hook, err := logrus_sentry.NewWithClientSentryHook(client, []logrus.Level{
    logrus.PanicLevel,
    logrus.FatalLevel,
    logrus.ErrorLevel,
  })

  if err == nil {
    log.Hooks.Add(hook)
  }

hook, err := NewWithClientSentryHook(client, []logrus.Level{
	logrus.ErrorLevel,
})


	if env == "production" {
		log.Info("Env is in production mode")
		log.SetLevel(log.ErrorLevel)
	} else {
		log.SetLevel(log.DebugLevel)
		log.Debug("Logging everything!!")
	}
}

// @BasePath /v1
// @title Hackerlog API
// @version 1.0
// @description This is the Hackerlog API for collecting editor stats

// @contact.name Deric Cain
// @contact.url https://dericcain.com
// @contact.email deric.cain@gmail.com

// @BasePath /v1
func main() {
	db := DbInit()
	defer db.Close()

	migrate(db)

	r := gin.Default()

	v1 := r.Group("/v1")

	AuthRoutes(v1.Group("/auth"))
	UserRoutes(v1.Group("/users"))
	UnitRoutes(v1.Group("/units"))
	CoreRoutes(v1.Group("/core"))

	// Setup Swagger docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(getPort())
}
