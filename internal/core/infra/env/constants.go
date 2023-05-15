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
	// SslModeDisable	don't care about security, and don't want to pay the overhead of encryption.
	SslModeDisable DBSslMode = "disable"

	// SslModeAllow don't care about security, but will pay the overhead of encryption if the server
	// insists on it.
	SslModeAllow DBSslMode = "allow"

	// SslModePrefer don't care about encryption, but wish to pay the overhead of encryption if the
	// server supports it.
	SslModePrefer DBSslMode = "prefer"

	// SslModeRequire want data to be encrypted, and accept the overhead. It trusts that the network
	// will make sure to always connect to the server requested.
	SslModeRequire DBSslMode = "require"

	// SslModeVerifyCa want data encrypted, and accept the overhead. It want to be sure that connect
	// to a server that it trusts.
	SslModeVerifyCa DBSslMode = "verify-ca"

	// SslModeVerifyFull want data encrypted, and accept the overhead. It want to be sure that
	// connect to a server that it trusts and that it is talking to the right one.
	SslModeVerifyFull DBSslMode = "verify-full"
)
