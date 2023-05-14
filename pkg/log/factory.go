package log

import "github.com/christian-gama/nutrai-api/internal/shared/domain/logger"

func MakeLogWithCaller(skip int) logger.Logger {
	return New(&Config{Caller: true, CallerSkip: skip, Stack: false})
}

func MakeLog() logger.Logger {
	return New(&Config{Caller: false, Stack: false})
}

func MakeLogWithStack() logger.Logger {
	return New(&Config{Caller: true, Stack: true})
}
