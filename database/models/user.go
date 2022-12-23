package models

import (
	"database/sql"
	"github.com/MyGoBB/MyGoBB/database"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"

	"github.com/MyGoBB/MyGoBB/utils/crypto"
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

// CreateUser creates a new user account and sends and activation email
func CreateUser(username string, password string, email string) *User {
	user := &User{
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// set the users password
	user.SetPassword(password)

	db, err := database.GetConnection()
	if err != nil {
		log.WithError(err).Error("Failed to get database connection")
		return nil
	}

	db.Create(user)
	return user
}

func (user *User) SetPassword(password string) {
	hashed, err := crypto.GenerateHashedPassowrd(password)
	if err == nil {
		user.Password = hashed
	}
}
