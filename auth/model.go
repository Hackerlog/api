package auth

import (
	"github.com/dericgw/blog-api/common"
	"github.com/dericgw/blog-api/users"
	"github.com/jinzhu/gorm"
)

// Auth These are the users' JWT's
type Auth struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	user      users.User
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

// Migrate Migrate the users table
func Migrate() {
	db := common.GetDb()

	db.AutoMigrate(&Auth{})
}
