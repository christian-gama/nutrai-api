package diet

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
	"github.com/christian-gama/nutrai-api/internal/diet/infra/validation"
	. "github.com/christian-gama/nutrai-api/internal/diet/module"
)

// Init is the function that initializes this module.
func Init() (*module.Module, func()) {
	return Module, func() {
		// Add the logic to initialize this module here:
		validation.Register()
	}
}
