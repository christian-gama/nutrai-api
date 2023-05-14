package env

const (
	// Production is the production string representation of the app environment.
	Production AppEnv = "prod"

	// Development is the development string representation of the app environment.
	Development AppEnv = "dev"

	// Test is the test string representation of the app environment.
	Test AppEnv = "test"
)

const (
	SslModeDisable    DBSslMode = "disable"
	SslModeAllow      DBSslMode = "allow"
	SslModePrefer     DBSslMode = "prefer"
	SslModeRequire    DBSslMode = "require"
	SslModeVerifyCa   DBSslMode = "verify-ca"
	SslModeVerifyFull DBSslMode = "verify-full"
)
