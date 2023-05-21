package event

import "github.com/christian-gama/nutrai-api/internal/core/domain/event"

var SaveException = event.New("exception", event.Save)
