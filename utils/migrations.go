package utils

import (
    "fmt"
    "go/build"
    "path/filepath"
    "os"

    "github.com/MyGoBB/MyGoBB/config"
    "github.com/MyGoBB/MyGoBB/models"
    "bitbucket.org/liamstask/goose/lib/goose"
    "github.com/spf13/viper"
    log "github.com/sirupsen/logrus"
)

var goose_conf *goose.DBConf

// GenerateGooseDbConf generates a new DBConf config for goose
func GenerateGooseDbConf() *goose.DBConf {
    if goose_conf != nil {
        return goose_conf
    }

    //pkg, _ := build.Import("github.com/MyGoBB/MyGoBB/mygobb", ".", build.FindOnly)
    db_host := viper.GetString(config.DatabaseHost)
    db_port := viper.GetInt(config.DatabasePort)
    db_name := viper.GetString(config.DatabaseName)
    db_username := viper.GetString(config.DatabaseUsername)
    db_password := viper.GetString(config.DatabasePassword)
    migrations_path := filepath.Clean("db/migrations")

    goose_conf = &goose.DBConf{
        MigrationsDir: migrations_path,
        Env: "development",
        Driver: goose.DBDriver{
            Name: "mysql",
            OpenStr: fmt.Sprintf("%s:%s@%s:%s/%s", db_username, db_password, db_host, db_port, db_name),
            Import: "github.com/go-sql-driver/mysql",
            Dialect: &goose.MySqlDialect{},
        },
    }

    return goose_conf
}

// GetMigrationInfo gets info regarding any pending migrations
func GetMigrationInfo(latest_db_version int64, migrations []*goose.Migration, err error) {
    goose_conf = GenerateGooseDbConf()
    db := models.GetDbSession()

    latest_db_version, _ = goose.GetMostRecentDBVersion(goose_conf.MigrationsDir)
    current_db_version, _ := goose.EnsureDBVersion(goose_conf, db.Db)
    migrations, _ = goose.CollectMigrations(goose_conf.MigrationsDir, current_db_version, latest_db_version)

    return latest_db_version, migrations, err
}

// RunMigrations runs the database migrations
func RunMigrations(version int64) error {
    goose_conf := GenerateGooseDbConf()
    return goose.RunMigrations(goose_conf, goose_conf.MigrationsDir, version)
}