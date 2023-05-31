package event

import "github.com/christian-gama/nutrai-api/internal/core/domain/event"

const Exception = "exception"

var SaveException = event.New(Exception, event.Save)
