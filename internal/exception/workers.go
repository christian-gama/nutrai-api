package exception

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/worker"
	"github.com/christian-gama/nutrai-api/internal/exception/app/consumer"
)

// Workers loads all workers for this module.
func Workers() {
	worker.Create(consumer.MakeSaveExceptionHandler().Handle, 1)
}
