package internal

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/router/routing"
	"github.com/christian-gama/nutrai-api/internal/user"
)

// Routes registers all routes from the application.
func Routes() {
	routing.Register("api",
		user.Routes(),
	)
}
