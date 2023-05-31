package event

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/event"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
)

const Exception = "exception"

var CatchException = event.New[exception.Exception](Exception, event.Error)
