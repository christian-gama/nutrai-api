package notify

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
	"github.com/christian-gama/nutrai-api/internal/core/infra/worker"
	. "github.com/christian-gama/nutrai-api/internal/core/module"
	"github.com/christian-gama/nutrai-api/internal/notify/app/consumer"
)

// Init is the function that initializes this module.
func Init() (*module.Module, func()) {
	return Module, func() {
		// Add the logic to initialize this module here:
		worker.Register(consumer.MakeSendWelcomeHandler().Handle, 2)
	}
}
