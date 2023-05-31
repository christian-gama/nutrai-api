package event

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/core/domain/event"
)

const User = "user"

var SaveUser = event.New[user.User](User, event.Save)
