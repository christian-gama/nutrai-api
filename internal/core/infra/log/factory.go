package log

import "github.com/christian-gama/nutrai-api/internal/core/domain/logger"

func MakeDefault() logger.Logger {
	return New(&Config{Caller: false, Stack: false})
}

func MakeWithCaller() logger.Logger {
	return New(&Config{Caller: true, CallerSkip: 1, Stack: false})
}

func MakeWithStack() logger.Logger {
	return New(&Config{Caller: true, Stack: true})
}
