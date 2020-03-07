package models

import (
	"database/sql"
	"fmt"

	"github.com/MyGoBB/MyGoBB/config"
	"github.com/go-gorp/gorp"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var db_map *gorp.DbMap

// GetDbSession gets the database session
func GetDbSession() *gorp.DbMap {
	if db_map != nil {
		return db_map
	}

	db_host := viper.GetString(config.DatabaseHost)
	db_port := viper.GetInt(config.DatabasePort)
	db_name := viper.GetString(config.DatabaseName)
	db_username := viper.GetString(config.DatabaseUsername)
	db_password := viper.GetString(config.DatabasePassword)

	/*var connectionString string

	if db_password != "" {
		connectionString = fmt.Sprintf("%s:%s@%s:%d/%s", db_username, db_password, db_host, db_port, db_name)
	} else {
		connectionString = fmt.Sprintf("%s@%s:%d/%s", db_username, db_host, db_port, db_name)
	}*/

	//db, err := sql.Open("mysql", connectionString)

	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s", db_username, db_password, db_host, db_port, db_name))
	db, err := sql.Open("postgres", fmt.Sprintf("username=%s password=%s dbname=%s host=%s port=%d, sslmode=disable", db_username, db_password, db_name, db_host, db_port))

	if err != nil {
		log.WithError(err).Error("Could not connect to database.")
		return nil
	}

	/*db_map = &gorp.DbMap{
		Db:      db,
		Dialect: gorp.MySQLDialect{},
	}*/

	db_map = &gorp.DbMap{
		Db:      db,
		Dialect: gorp.PostgresDialect{},
	}

	db_map.AddTableWithName(User{}, "users").SetKeys(true, "Id")

	return db_map
}
