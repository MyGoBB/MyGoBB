package services

import (
	"time"

	"github.com/MyGoBB/MyGoBB/database"
	"github.com/MyGoBB/MyGoBB/database/models"

	log "github.com/sirupsen/logrus"
)

// CreateUser creates a new user account and sends and activation email
func CreateUser(username string, password string, email string) *models.User {
	user := &models.User{
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
