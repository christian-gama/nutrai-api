package log

import (
	"os"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerImpl struct {
	base *zap.SugaredLogger
}

// Config is the configuration for the logger.
type Config struct {
	Stack      bool
	Caller     bool
	CallerSkip int
}

// New creates a new logger for the application. It uses the
// zap.NewProductionEncoderConfig to create an encoder config that
// will be used to create the encoder for the logger. The encoder
// is then used to create a zapcore.Core which is then used to
// create a new logger. The logger is configured to log at the
// InfoLevel and will also log stack traces for errors.
//
// The function returns a logger that has been converted to a
// sugared logger.
func New(config *Config) logger.Logger {
	if config == nil {
		config = &Config{}
	}

	if config.CallerSkip <= 0 {
		config.CallerSkip = 1
	}

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncodeDuration = zapcore.StringDurationEncoder
	cfg.EncodeName = zapcore.FullNameEncoder
	cfg.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(strings.Split(caller.TrimmedPath(), ":")[0])
	}

	encoder := zapcore.NewConsoleEncoder(cfg)
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)

	options := []zap.Option{}
	if config.Stack {
		options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	}
	if config.Caller {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(config.CallerSkip))
	}

	logger := zap.New(core, options...)

	return &loggerImpl{base: logger.Sugar()}
}

// Debug implements logger.Logger.
func (l *loggerImpl) Debug(args ...any) {
	l.base.Debug(args...)
}

// Debugf implements logger.Logger.
func (l *loggerImpl) Debugf(format string, args ...any) {
	l.base.Debugf(format, args...)
}

// Error implements logger.Logger.
func (l *loggerImpl) Error(args ...any) {
	l.base.Error(args...)
}

// Errorf implements logger.Logger.
func (l *loggerImpl) Errorf(format string, args ...any) {
	l.base.Errorf(format, args...)
}

// Fatal implements logger.Logger.
func (l *loggerImpl) Fatal(args ...any) {
	l.base.Fatal(args...)
}

// Fatalf implements logger.Logger.
func (l *loggerImpl) Fatalf(format string, args ...any) {
	l.base.Fatalf(format, args...)
}

// Info implements logger.Logger.
func (l *loggerImpl) Info(args ...any) {
	l.base.Info(args...)
}

// Infof implements logger.Logger.
func (l *loggerImpl) Infof(format string, args ...any) {
	l.base.Infof(format, args...)
}

// Warn implements logger.Logger.
func (l *loggerImpl) Warn(args ...any) {
	l.base.Warn(args...)
}

// Warnf implements logger.Logger.
func (l *loggerImpl) Warnf(format string, args ...any) {
	l.base.Warnf(format, args...)
}

// With implements logger.Logger.
func (l *loggerImpl) With(args ...any) logger.Logger {
	return &loggerImpl{
		base: l.base.With(args...),
	}
}

// Panic implements logger.Logger.
func (l *loggerImpl) Panic(args ...any) {
	l.base.Panic(args...)
}

// Panicf implements logger.Logger.
func (l *loggerImpl) Panicf(format string, args ...any) {
	l.base.Panicf(format, args...)
}
