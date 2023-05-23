package env

import (
	"fmt"

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
			return fmt.Errorf(
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
			return fmt.Errorf(
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
			return fmt.Errorf(
				"Invalid env variable: '%s'. Must be one of: %v",
				Config.LogLevel,
				validLogLevels,
			)
		}
		return nil
	}
}
