package event

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/event"
)

const Exception = "exception"

var CatchException = event.New(Exception, event.Error)
