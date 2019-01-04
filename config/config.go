package config

import (
    "github.com/spf13/viper"
)

// LoadConfig loads the configuration from a specified file
func LoadConfig(path string) error {
    if path != "" {
        viper.SetConfigFile(path)
    } else {
        viper.AddConfigPath("./")
		viper.SetConfigType("yaml")
        viper.SetConfigName("config")
    }

    // Find and read the config file
    if err := viper.ReadInConfig(); err != nil {
        return err
    }

    return nil
}

// SaveConfig stores the configuration to a specified file
func SaveConfig(path string) error {
    // TODO: Implement

    return nil
}

// SetDefaults sets the default options
func SetDefaults() {
    viper.SetDefault(Debug, false)
    viper.SetDefault(APIHost, "0.0.0.0")
    viper.SetDefault(APIPort, 8080)
    viper.SetDefault(SSLEnabled, false)
    viper.SetDefault(SSLGenerateLetsencrypt, false)
    viper.SetDefault(UploadsMaximumSize, 100000)
    viper.SetDefault(LogPath, "./logs")
    viper.SetDefault(LogLevel, "info")
    viper.SetDefault(LogDeleteAfterDays, "30")
}

// ContainsAuthKey checks wether the config contains a specified authentication key
func ContainsAuthKey(key string) bool {
    for _, k := range viper.GetStringSlice(AuthKey) {
        if k == key {
            return true
        }
    }
    return false
}