package log

import "github.com/christian-gama/nutrai-api/internal/core/domain/logger"

func MakeDefault() logger.Logger {
	return New(&Config{Caller: false, Stack: false})
}

func MakeWithCaller(skip ...int) logger.Logger {
	s := 1
	if len(skip) > 0 {
		s = skip[0]
	}

	return New(&Config{Caller: true, CallerSkip: s, Stack: false})
}

func MakeWithStack() logger.Logger {
	return New(&Config{Caller: true, Stack: true})
}
