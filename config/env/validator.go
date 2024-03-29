package env

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/christian-gama/nutrai-api/pkg/slice"
)

type validator func() error

func validateDBSslMode() validator {
	return func() error {
		validSslModes := []DBSslMode{
			SslModeDisable,
			SslModeAllow,
			SslModePrefer,
			SslModeRequire,
			SslModeVerifyCa,
			SslModeVerifyFull,
		}

		if !slice.Contains(validSslModes, DB.SslMode) {
			return errors.InternalServerError(
				"Invalid env variable: '%s'. Must be one of: %v",
				DB.SslMode,
				validSslModes,
			)
		}
		return nil
	}
}

func validateAppEnv() validator {
	return func() error {
		validEnvs := []AppEnv{Production, Development, Test}

		if !slice.Contains(validEnvs, App.Env) {
			return errors.InternalServerError(
				"Invalid env variable: '%s'. Must be one of: %v",
				App.Env,
				validEnvs,
			)
		}
		return nil
	}
}

func validateConfigLogLevel() validator {
	return func() error {
		validLogLevels := []ConfigLogLevel{
			LogLevelInfo,
			LogLevelWarn,
			LogLevelError,
			LogLevelDebug,
			LogLevelPanic,
		}

		if !slice.Contains(validLogLevels, Config.LogLevel) {
			return errors.InternalServerError(
				"Invalid env variable: '%s'. Must be one of: %v",
				Config.LogLevel,
				validLogLevels,
			)
		}
		return nil
	}
}

func validateMailerProvider() validator {
	return func() error {
		validProviders := []MailerProvider{
			MailerProviderMailtrap,
			MailerProviderSendgrid,
		}

		if !slice.Contains(validProviders, Mailer.Provider) {
			return errors.InternalServerError(
				"Invalid env variable: '%s'. Must be one of: %v",
				Mailer.Provider,
				validProviders,
			)
		}
		return nil
	}
}
