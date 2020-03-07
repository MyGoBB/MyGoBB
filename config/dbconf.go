package config

import "database/sql"

// DBConfig contains the database configuration
type DBConfig struct {
	MigrationsDir string
	Driver DBDriver
}

// DBDriver contains information about the database driver
type DBDriver struct {
	Name string
	OpenStr string
	Db *sql.DB
}
