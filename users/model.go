package users

import (
	"errors"
	"time"

	"github.com/hackerlog/api/common"
	"github.com/hackerlog/api/units"
	"github.com/pborman/uuid"
)

// User This is the user model that will hold all of the users
type User struct {
	ID                 uint         `json:"id" gorm:"primary_key"`
	Email              string       `json:"email" gorm:"type:varchar(100);unique_index" binding:"required"`
	Name               string       `json:"name" binding:"required"`
	Password           string       `json:"-" binding:"required"`
	EditorToken        string       `json:"editor_token" gorm:"index"`
	PasswordResetToken string       `json:"-"`
	Units              []units.Unit `json:"units"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
	DeletedAt          *time.Time   `json:"-" sql:"index"`
}

// BeforeCreate We want to hash the users password
func (u *User) BeforeCreate() (err error) {
	hashedPassword, err := common.HashPassword(u.Password)

	if err != nil {
		err = errors.New("Hashing user password failed")
	}

	u.Password = hashedPassword
	u.EditorToken = uuid.New()

	return nil
}
