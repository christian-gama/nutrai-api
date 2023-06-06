package metrics

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
	"github.com/christian-gama/nutrai-api/internal/metrics/api/http/routes"
	"github.com/christian-gama/nutrai-api/internal/metrics/infra/metrics"
	"github.com/christian-gama/nutrai-api/internal/metrics/infra/sql/hook"
	. "github.com/christian-gama/nutrai-api/internal/metrics/module"
)

// Init is the function that initializes this module.
func Init() (*module.Module, func()) {
	return Module, func() {
		// Add the logic to initialize this module here:
		metrics.Register()
		hook.Register()
		routes.Register()
	}
}
