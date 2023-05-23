package log

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"go.uber.org/zap/zapcore"
)

var level = map[env.ConfigLogLevel]zapcore.Level{
	env.LogLevelDebug: zapcore.DebugLevel,
	env.LogLevelInfo:  zapcore.InfoLevel,
	env.LogLevelWarn:  zapcore.WarnLevel,
	env.LogLevelError: zapcore.ErrorLevel,
	env.LogLevelPanic: zapcore.PanicLevel,
}
