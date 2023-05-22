package diet

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/internal/diet/module"
)

// Init is the function that initializes this module.
func Init(log logger.Logger) {
	module.Module.Init(log, func() {
		// Add the logic to initialize this module here:
	})
}
