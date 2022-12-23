package models

import (
	"database/sql"
	"time"

	"github.com/MyGoBB/MyGoBB/utils/crypto"
	"gorm.io/gorm"
)

// User model
type User struct {
	Id               uint64 `gorm:"primaryKey"`
	Username         string `gorm:"unique"`
	Password         string
	Email            string `gorm:"unique"`
	Avatar           sql.NullString
	Signature        sql.NullString
	EmailConfirmedAt sql.NullTime
	LastSeen         sql.NullTime
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (user *User) SetPassword(password string) {
	hashed, err := crypto.GenerateHashedPassowrd(password)
	if err == nil {
		user.Password = hashed
	}
}
