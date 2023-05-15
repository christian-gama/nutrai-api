package internal

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/router/routing"
	"github.com/christian-gama/nutrai-api/internal/user"
)

// Routing registers all routes from the application.
func Routing() {
	routing.Register("api",
		user.Routes(),
	)
}
