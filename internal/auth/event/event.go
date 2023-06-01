package event

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/event"
)

const User = "user"

var SaveUser = event.New(User, event.Save)
