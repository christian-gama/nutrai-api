package exception

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/internal/core/infra/worker"
	"github.com/christian-gama/nutrai-api/internal/exception/api/http"
	"github.com/christian-gama/nutrai-api/internal/exception/app/consumer"
	"github.com/christian-gama/nutrai-api/internal/exception/module"
)

// Init is the function that initializes this module.
func Init(log logger.Logger) {
	module.Module.Init(log, func() {
		// Add the logic to initialize this module here:
		worker.Create(consumer.MakeSaveExceptionHandler().Handle, 1)
		http.Routes()
	})
}
