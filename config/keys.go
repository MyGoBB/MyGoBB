package config

const (
	// Debug is a boolean value that enables debug mode
	Debug = "debug"

	// APIHost is a string containing the interface ip address
	// on what the api should listen on
	APIHost = "api.host"

	// APIPort is an integer containing the port the api should
	// listen on
	APIPort = "api.port"

	// SSLEnabled is a boolean that states whether ssl should be enabled or not
	SSLEnabled = "api.ssl.enabled"

	// SSLGenerateLetsencrypt is a boolean that enables automatic SSL certificate
	// creation with letsencrypt
	SSLGenerateLetsencrypt = "api.ssl.letsencrypt"

	// SSLCertificate is a string containing the location of
	// a ssl certificate to use
	SSLCertificate = "api.ssl.cert"

	// SSLKey is a string containing the location of the key
	// for the ssl certificate
	SSLKey = "api.ssl.key"

	// UploadsMaximumSize is an integer that sets the maximum size for
	// file uploads through the api in Kilobytes
	UploadsMaximumSize = "api.uploads.maximumSize"

	// LogPath is a string containing the path where logfiles should be
	// stored
	LogPath = "log.path"

	// LogLevel is a string containing the log level
	LogLevel = "log.level"

	// LogDeleteAfterDays is an integer containing the amounts of days
	// logs should be stored. They will be deleted after. If set to 0
	// logs will be stored indefinitely.
	LogDeleteAfterDays = "log.deleteAfterDays"

	// AuthKey contains a key that will be replaced by something better
	AuthKey = "authKey"

	// DatabaseType is the type of database we want to use. Only postgresql and mysql supported at the moment
	DatabaseType = "database.type"

	// DatabaseHost is the host for the database
	DatabaseHost = "database.host"

	// DatabasePort is the port used when connecting to the database
	DatabasePort = "database.port"

	// DatabaseName is the name of the database
	DatabaseName = "database.name"

	// DatabaseUsername is the username to use for connecting to the database
	DatabaseUsername = "database.username"

	// DatabasePassword is the password to use when connecting to the database
	DatabasePassword = "database.password"
)
