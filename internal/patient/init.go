package patient

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/internal/patient/api/http"
	"github.com/christian-gama/nutrai-api/internal/patient/module"
)

// Init is the function that initializes this module.
func Init(log logger.Logger) {
	module.Module.Init(log, func() {
		// Add the logic to initialize this module here:
		http.Routes()
	})
}
