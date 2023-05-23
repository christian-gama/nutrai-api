package exception

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
	"github.com/christian-gama/nutrai-api/internal/core/infra/worker"
	"github.com/christian-gama/nutrai-api/internal/exception/api/http"
	"github.com/christian-gama/nutrai-api/internal/exception/app/consumer"
	. "github.com/christian-gama/nutrai-api/internal/exception/module"
)

// Init is the function that initializes this module.
func Init() (*module.Module, func()) {
	return Module, func() {
		worker.Create(consumer.MakeSaveExceptionHandler().Handle, 1)
		http.Routes()
	}
}
