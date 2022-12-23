package database

import (
	"errors"

	"github.com/MyGoBB/MyGoBB/database/models"
)

func RunMigrations() error {
	db, err := GetConnection()
	if err != nil {
		return errors.New("failed to get database connection")
	}

	// migrate schema
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return errors.New("could not migrate users schema")
	}

	if err := db.AutoMigrate(&models.Forum{}); err != nil {
		return errors.New("could not migrate forums schema")
	}

	if err := db.AutoMigrate(&models.Topic{}); err != nil {
		return errors.New("could not migrate topics schema")
	}

	return nil
}
