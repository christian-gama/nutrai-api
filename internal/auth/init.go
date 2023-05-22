package auth

import (
	"github.com/christian-gama/nutrai-api/internal/auth/api/http"
	"github.com/christian-gama/nutrai-api/internal/auth/module"
	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
)

// Init is the function that initializes this module.
func Init(log logger.Logger) {
	module.Module.Init(log, func() {
		// Add the logic to initialize this module here:
		http.Routes()
	})
}
