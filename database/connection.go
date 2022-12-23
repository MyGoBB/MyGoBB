package database

import (
	"errors"
	"fmt"
	"github.com/MyGoBB/MyGoBB/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

var connection *gorm.DB

func GetConnection() (*gorm.DB, error) {
	if connection != nil {
		return connection, nil
	}

	dbType := viper.GetString(config.DatabaseType)

	dbHost := viper.GetString(config.DatabaseHost)
	dbPort := viper.GetInt(config.DatabasePort)
	dbName := viper.GetString(config.DatabaseName)
	dbUsername := viper.GetString(config.DatabaseUsername)
	dbPassword := viper.GetString(config.DatabasePassword)

	if dbType == "" {
		return nil, errors.New("invalid database type or database type is missing")
	}

	if strings.EqualFold(dbType, "mysql") {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			connection = db
			return connection, nil
		}

		return db, nil
	} else if strings.EqualFold(dbType, "postgresql") {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUsername, dbPassword, dbName, dbPort)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			connection = db
			return connection, nil
		}

		return db, nil
	} else {
		return nil, errors.New("invalid database type or database type is missing")
	}
}
