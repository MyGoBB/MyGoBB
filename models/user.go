package models

import (
	"database/sql"
	"time"

	"github.com/MyGoBB/MyGoBB/utils/crypto"
	"github.com/lib/pq"
)

// User model
type User struct {
	Id               int64          `db:"id"`
	Username         string         `db:"username"`
	Password         string         `db:"password"`
	Email            string         `db:"email"`
	Avatar           sql.NullString `db:"avatar"`
	Signature        sql.NullString `db:"signature"`
	EmailConfirmedAt pq.NullTime    `db:"email_confirmed_at"`
	CreatedOn        time.Time      `db:"created_on"`
	UpdatedAt		 time.Time		`db:"created_at"`
	LastSeen         pq.NullTime    `db:"last_seen"`
}

// CreateUser creates a new user account and sends and activation email
func CreateUser(username string, password string, email string) *User {
	user := &User{
		Username:  username,
		Email:     email,
		CreatedOn: time.Now(),
		UpdatedAt: time.Now(),
	}

	// set the users password
	user.SetPassword(password)
	return user
}

func (user *User) SetPassword(password string) {
	hashed, err := crypto.GenerateHashedPassowrd(password)
	if err == nil {
		user.Password = hashed
	}
}
