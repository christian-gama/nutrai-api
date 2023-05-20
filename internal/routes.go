package internal

import (
	"github.com/christian-gama/nutrai-api/internal/auth"
	"github.com/christian-gama/nutrai-api/internal/core/infra/router/routing"
	"github.com/christian-gama/nutrai-api/internal/patient"
)

// Routing registers all routes from the application.
func Routing() {
	routing.Register("api",
		patient.Routes(),
		auth.Routes(),
	)
}
