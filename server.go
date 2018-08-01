package main

import (
	"os"

	_ "github.com/hackerlog/api/docs"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	raven "github.com/getsentry/raven-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/ulule/limiter"
	mgin "github.com/ulule/limiter/drivers/middleware/gin"
	memory "github.com/ulule/limiter/drivers/store/memory"
)

var (
	env      = os.Getenv("APP_ENV")
	xHeader  = "X-Hackerlog-EditorToken"
	xpHeader = "X-Hackerlog-PurgeToken"
)

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

	raven.SetDSN(os.Getenv("SENTRY_DSN"))
	raven.SetEnvironment(env)

	if env != "local" {
		log.SetLevel(log.ErrorLevel)
	} else {
		log.SetLevel(log.DebugLevel)
		log.Debug("Logging everything!!")
	}
}

// @BasePath /v1
// @title Hackerlog API
// @version v0.1
// @description This is the Hackerlog API
// @contact.name Deric Cain
// @contact.email deric.cain@gmail.com
// @BasePath /v1
func main() {
	db := DbInit()
	defer db.Close()

	migrate(db)

	if env != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(cors.Default())

	// Setup rate limiting
	rate, err := limiter.NewRateFromFormatted("500-m")

	if err != nil {
		log.Error("Rate limiting crapped out", err)
	}

	store := memory.NewStore()
	rateLimiting := mgin.NewMiddleware(limiter.New(store, rate))

	r.ForwardedByClientIP = true
	r.Use(rateLimiting)

	// Setup Sentry logging
	r.Use(sentry.Recovery(raven.DefaultClient, false))

	// Wrap all routes in v1/ URL
	v1 := r.Group("/v1")

	AuthRoutes(v1.Group("/auth"))
	UserRoutes(v1.Group("/users"))
	UnitRoutes(v1.Group("/units"))
	CoreRoutes(v1.Group("/core"))
	MailingListRoutes(v1.Group("/mailing-list"))

	// Setup Swagger docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(getPort())
}
