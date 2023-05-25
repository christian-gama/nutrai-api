package log

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/pkg/slice"
)

func MakeDefault() logger.Logger {
	return New(&Config{Caller: false, Stack: false})
}

func MakeWithCaller(skip ...int) logger.Logger {
	return New(
		&Config{
			Caller:     true,
			CallerSkip: slice.FirstElementOrDefault(skip, 1),
			Stack:      false,
		},
	)
}

func MakeWithStack() logger.Logger {
	return New(&Config{Caller: true, Stack: true})
}
