package utils

import (
	"fmt"
	"path/filepath"

	"github.com/MyGoBB/MyGoBB/config"
	"github.com/MyGoBB/MyGoBB/models"
	//"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/pressly/goose"
	//log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//var goose_conf *goose.DBConf
var dbconf *config.DBConfig

// GnerateDbConf generates a new DBConfig for migrations
func GenerateDbConf() *config.DBConfig {
	if dbconf != nil {
		return dbconf
	}

	db := models.GetDbSession()

	//pkg, _ := build.Import("github.com/MyGoBB/MyGoBB/mygobb", ".", build.FindOnly)
	db_host := viper.GetString(config.DatabaseHost)
	db_port := viper.GetInt(config.DatabasePort)
	db_name := viper.GetString(config.DatabaseName)
	db_username := viper.GetString(config.DatabaseUsername)
	db_password := viper.GetString(config.DatabasePassword)
	//migrations_path := filepath.Clean("db/migrations")
	//migrations_path := filepath.Join(pkg.SrcRoot, pkg.ImportPath, "../db/migrations")
	path, _ := GetExecutableDir()
	migrations_path := filepath.Join(path, "/db/migrations")

	dbconf = &config.DBConfig{
		MigrationsDir: migrations_path,
		Driver: config.DBDriver{
			Name: "postgres",
			OpenStr: fmt.Sprintf("username=%s password=%s dbname=%s host=%s port=%d, sslmode=disable", db_username, db_password, db_name, db_host, db_port),
			Db: db.Db,
		},
	}

	return dbconf
}

// GenerateGooseDbConf generates a new DBConf config for goose
/*func GenerateGooseDbConf() *goose.DBConf {
	if goose_conf != nil {
		return goose_conf
	}

	pkg, _ := build.Import("github.com/MyGoBB/MyGoBB/mygobb", ".", build.FindOnly)
	db_host := viper.GetString(config.DatabaseHost)
	db_port := viper.GetInt(config.DatabasePort)
	db_name := viper.GetString(config.DatabaseName)
	db_username := viper.GetString(config.DatabaseUsername)
	db_password := viper.GetString(config.DatabasePassword)
	migrations_path := filepath.Join(pkg.SrcRoot, pkg.ImportPath, "../db/migrations")

	goose_conf = &goose.DBConf{
		MigrationsDir: migrations_path,
		Env:           "development",
		Driver: goose.DBDriver{
			Name:    "mysql",
			OpenStr: fmt.Sprintf("%s:%s@%s:%d/%s", db_username, db_password, db_host, db_port, db_name),
			Import:  "github.com/go-sql-driver/mysql",
			Dialect: &goose.MySqlDialect{},
		},
	}

	goose_conf = &goose.DBConf{
		MigrationsDir: migrations_path,
		Env:           "development",
		Driver: goose.DBDriver{
			Name:    "postgres",
			OpenStr: fmt.Sprintf("username=%s password=%s dbname=%s host=%s port=%d, sslmode=disable", db_username, db_password, db_name, db_host, db_port),
			Import:  "github.com/lib/pq",
			Dialect: &goose.PostgresDialect{},
		},
	}

	return goose_conf
}*/

// GetMigrationInfo gets info regarding any pending migrations
func GetMigrationInfo() (latest_db_version int64, migrations []*goose.Migration, err error) {
	gooseConf := GenerateDbConf()

	latest_db_version, err = goose.GetDBVersion(gooseConf.Driver.Db)
	current, err := goose.EnsureDBVersion(gooseConf.Driver.Db)
	//latest_db_version, err = goose.GetMostRecentDBVersion(goose_conf.MigrationsDir)
	migrations, err = goose.CollectMigrations(gooseConf.MigrationsDir, current, latest_db_version)

	return latest_db_version, migrations, err
}

/*func GetMigrationInfo() (latest_db_version int64, migrations []*goose.Migration, err error) {
	goose_conf = GenerateGooseDbConf()
	db := models.GetDbSession()

	latest_db_version, err = goose.GetMostRecentDBVersion(goose_conf.MigrationsDir)
	current_db_version, err := goose.EnsureDBVersion(goose_conf, db.Db)
	migrations, err = goose.CollectMigrations(goose_conf.MigrationsDir, current_db_version, latest_db_version)

	return latest_db_version, migrations, err
}*/

// RunMigrations runs the database migrations
func RunMigrations(version int64) error {
	dbconf = GenerateDbConf()

	return goose.Up(dbconf.Driver.Db, dbconf.MigrationsDir)
}

/*func RunMigrations(version int64) error {
	goose_conf := GenerateGooseDbConf()
	db := models.GetDbSession()

	return goose.RunMigrationsOnDb(goose_conf, goose_conf.MigrationsDir, version, db.Db)
}*/



