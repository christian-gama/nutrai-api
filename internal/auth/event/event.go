package event

import "github.com/christian-gama/nutrai-api/internal/core/domain/event"

var SaveUser = event.New("user", event.Save)
