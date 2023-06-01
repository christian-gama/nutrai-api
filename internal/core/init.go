package core

import (
	"github.com/christian-gama/nutrai-api/internal/core/api/http/routes"
	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
	. "github.com/christian-gama/nutrai-api/internal/core/module"
)

// Init is the function that initializes this module.
func Init() (*module.Module, func()) {
	return Module, func() {
		// Add the logic to initialize this module here:
		routes.Register()
	}
}
