package log

import (
	"log"
	"os"
	"strings"

	"github.com/christian-gama/nutrai-api/config/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a new logger for the application. It uses the zap.NewProductionEncoderConfig to
// create an encoder config that will be used to create the encoder for the logger. The encoder is
// then used to create a zapcore.Core which is then used to create a new logger. The logger is
// configured to log at the InfoLevel and will also log stack traces for errors.
func New() *zap.SugaredLogger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncodeDuration = zapcore.StringDurationEncoder
	cfg.EncodeName = zapcore.FullNameEncoder
	cfg.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(strings.Split(caller.TrimmedPath(), ":")[0])
	}

	encoder := zapcore.NewConsoleEncoder(cfg)
	if env.Config == nil {
		log.Panic("could not create logger - did you load the environment variables first?")
	}

	if _, ok := level[env.Config.LogLevel]; !ok {
		log.Panicf("could not find log level %s", env.Config.LogLevel)
	}
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level[env.Config.LogLevel])

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return logger.Sugar()
}

var SugaredLogger *zap.SugaredLogger

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...any) {
	SugaredLogger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...any) {
	SugaredLogger.Debugf(template, args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...any) {
	SugaredLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...any) {
	SugaredLogger.Infof(template, args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...any) {
	SugaredLogger.Warn(args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...any) {
	SugaredLogger.Warnf(template, args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...any) {
	SugaredLogger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...any) {
	SugaredLogger.Errorf(template, args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...any) {
	SugaredLogger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...any) {
	SugaredLogger.Panicf(template, args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...any) {
	SugaredLogger.Fatal(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...any) {
	SugaredLogger.Fatalf(template, args...)
}

// With adds a variadic number of fields to the logging context. It accepts a mix of strongly-typed
// Field objects and loosely-typed key-value pairs. When processing pairs, the first element of the
// pair is used as the field key and the second as the field value.
//
// For example,
//
//	sugaredLogger.With(
//	  "hello", "world",
//	  "failure", errors.New("oh no"),
//	  Stack(),
//	  "count", 42,
//	  "user", User{Name: "alice"},
//
// )
//
// is the equivalent of
//
// unsugared.With(
//
//	String("hello", "world"),
//	String("failure", "oh no"),
//	Stack(),
//	Int("count", 42),
//	Object("user", User{Name: "alice"}),
//
// )
//
// Note that the keys in key-value pairs should be strings. In development, passing a non-string key
// panics. In production, the logger is more forgiving: a separate error is logged, but the
// key-value pair is skipped and execution continues. Passing an orphaned key triggers similar
// behavior: panics in development and errors in production.
func With(args ...any) *zap.SugaredLogger {
	return SugaredLogger.With(args...)
}

// WithOptions clones the current logger, applies the supplied Options, and returns the resulting
// logger. It's safe to use concurrently.
func WithOptions(opts ...zap.Option) *zap.SugaredLogger {
	return SugaredLogger.WithOptions(opts...)
}

// Named adds a sub-scope to the logger's name. See Logger.Named for details.
func Named(s string) *zap.SugaredLogger {
	return SugaredLogger.Named(s)
}

// Desugar unwraps the SugaredLogger, exposing the original Logger. Desugaring is idempotent.
func Desugar() *zap.Logger {
	return SugaredLogger.Desugar()
}

// Sync flushes any buffered log entries.
func Sync() {
	SugaredLogger.Sync()
}

// Loading logs a message at the Info level.
func Loading(template string, args ...any) {
	SugaredLogger.Info(LoadingColor(template, args...))
}
