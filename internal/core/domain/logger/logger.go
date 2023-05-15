package logger

type Logger interface {
	// Debug logs a message at level Debug on the standard logger.
	Debug(args ...any)

	// Debugf logs a message at level Debug on the standard logger.
	Debugf(format string, args ...any)

	// Info logs a message at level Info on the standard logger.
	Info(args ...any)

	// Infof logs a message at level Info on the standard logger.
	Infof(format string, args ...any)

	// Warn logs a message at level Warn on the standard logger.
	Warn(args ...any)

	// Warnf logs a message at level Warn on the standard logger.
	Warnf(format string, args ...any)

	// Error logs a message at level Error on the standard logger.
	Error(args ...any)

	// Errorf logs a message at level Error on the standard logger.
	Errorf(format string, args ...any)

	// Fatal logs a message at level Fatal on the standard logger.
	Fatal(args ...any)

	// Fatalf logs a message at level Fatal on the standard logger.
	Fatalf(format string, args ...any)

	// Panic logs a message at level Panic on the standard logger.
	Panic(args ...any)

	// Panicf logs a message at level Panic on the standard logger.
	Panicf(format string, args ...any)

	// With returns a new Logger instance with the specified key/value pairs appended to its
	// context.
	With(args ...any) Logger
}
