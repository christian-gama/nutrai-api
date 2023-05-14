package env

const (
	// EnvProduction is the production string representation of the app environment.
	EnvProduction AppEnv = "prod"

	// EnvDevelopment is the development string representation of the app environment.
	EnvDevelopment AppEnv = "dev"

	// EnvTest is the test string representation of the app environment.
	EnvTest AppEnv = "test"
)

const (
	SslModeDisable    DBSslMode = "disable"
	SslModeAllow      DBSslMode = "allow"
	SslModePrefer     DBSslMode = "prefer"
	SslModeRequire    DBSslMode = "require"
	SslModeVerifyCa   DBSslMode = "verify-ca"
	SslModeVerifyFull DBSslMode = "verify-full"
)
