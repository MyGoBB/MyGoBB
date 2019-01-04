package command

import (
    "path/filepath"
    "strconv"

    "github.com/spf13/viper"

    "github.com/MyGoBB/MyGoBB/config"
    "github.com/MyGoBB/MyGoBB/constants"
    "github.com/MyGoBB/MyGoBB/utils"
    "github.com/MyGoBB/MyGoBB/models"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)


// RootCommand is the root command of mygobb
var RootCommand = &cobra.Command{
    Use: "mygobb",
    Short: "MyGoBB is a next generation simple and fast bulletin board written in Go.",
    Long: "MyGoBB is a next generation simple and fast bulletin board written in Go.",
    Run: run,
    Version: constants.Version,
}

var configPath string

func init() {
    RootCommand.Flags().StringVarP(&configPath, "config", "c", "./config.yml", "Allows to set the path of the configuration file.")
}

// Execute registers the RootCommand
func Execute() {
    RootCommand.Execute()
}

func run(cmd *cobra.Command, args []string) {
    utils.InitLogging()
    log.Info("Loading configuration...")
    if err := config.LoadConfig(configPath); err != nil {
        log.WithError(err).Fatal("Could not locate a suitable config.yml file. Aborting startup.")
        log.Exit(1)
    }
    utils.ConfigureLogging()

    log.Info("MyGoBB v" + constants.Version)
    log.Info()
    log.Info("Configuration loaded successfully.")

    log.Info("Loading database...")
    // TODO: load and connect to a mysql database
    latest_db_version, migrations, err := utils.GetMigrationInfo()
    if len(migrations) != 0 {
        log.Info("Running database migrations:")
        err = utils.RunMigrations(latest_db_version)
        if err != nil {
            log.WithError(err).Error("Could not run migrations")
        }
        log.Info("Database migration successful")
    }

    log.Info("Starting API server...")
    // TODO: implement an api server

    log.Info("Starting Frontend server...")
    // TODO: implement a frontend server
}