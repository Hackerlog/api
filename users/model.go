package users

import (
	"errors"

	"github.com/dericgw/blog-api/common"
	"github.com/jinzhu/gorm"
)

// User This is the user model that will hold all of the users
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(100);unique_index" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// BeforeCreate We want to hash the users password
func (u *User) BeforeCreate() (err error) {
	hashedPassword, err := common.HashPassword(u.Password)

	if err != nil {
		err = errors.New("Hashing user password failed")
	}

	u.Password = hashedPassword

	return nil
}

// Migrate Migrate the users table
func Migrate() {
	db := common.GetDb()

	db.AutoMigrate(&User{})
}
