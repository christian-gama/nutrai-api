package event

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/event"
)

const Exception = "exception"

var Recovery = event.New(Exception, event.Error)
