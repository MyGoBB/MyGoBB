package models

import (
    "fmt"
    "database/sql"

    "github.com/MyGoBB/MyGoBB/config"
    "github.com/go-gorp/gorp"
    _ "github.com/go-sql-driver/mysql"
    "github.com/spf13/viper"
    log "github.com/sirupsen/logrus"
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

    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%s/%s", db_username, db_password, db_host, db_port, db_name))

    if err != nil {
        log.WithError(err).Error("Could not connect to database.")
        return nil
    }

    db_map = &gorp.DbMap{
        Db: db,
        Dialect: gorp.MySQLDialect{},
    }

    return db_map
}