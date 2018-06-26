package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/dericgw/blog-api/auth"
	"github.com/dericgw/blog-api/common"
	"github.com/dericgw/blog-api/users"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var env = os.Getenv("APP_ENV")

func migrate() {
	users.Migrate()
	auth.Migrate()
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		return ":8000"
	}
	return ":" + port
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Error("No .env file found", err)
	}

	if env == "production" {
		log.SetLevel(log.ErrorLevel)
	}
}

func main() {
	db := common.Init()
	defer db.Close()

	migrate()

	r := gin.Default()

	v1 := r.Group("/api")
	auth.Routes(v1.Group("/auth"))
	users.Routes(v1.Group("/users"))

	r.Run(getPort())
}
