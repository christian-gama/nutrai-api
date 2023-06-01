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
	// Http is the http string representation of the app protocol.
	Http AppProtocol = "http"

	// Https is the https string representation of the app protocol.
	Https AppProtocol = "https"
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

const (
	// LogLevelInfo is the info string representation of the log level.
	LogLevelInfo ConfigLogLevel = "info"

	// LogLevelWarn is the warn string representation of the log level.
	LogLevelWarn ConfigLogLevel = "warn"

	// LogLevelError is the error string representation of the log level.
	LogLevelError ConfigLogLevel = "error"

	// LogLevelDebug is the debug string representation of the log level.
	LogLevelDebug ConfigLogLevel = "debug"

	// LogLevelPanic is the panic string representation of the log level.
	LogLevelPanic ConfigLogLevel = "panic"
)

const (
	MailerProviderSendgrid MailerProvider = "sendgrid"
	MailerProviderMailtrap MailerProvider = "mailtrap"
)
