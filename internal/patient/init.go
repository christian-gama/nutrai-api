package patient

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
	"github.com/christian-gama/nutrai-api/internal/patient/api/http"
	"github.com/christian-gama/nutrai-api/internal/patient/infra/validation"
	. "github.com/christian-gama/nutrai-api/internal/patient/module"
)

// Init is the function that initializes this module.
func Init() (*module.Module, func()) {
	return Module, func() {
		// Add the logic to initialize this module here:
		validation.Register()
		http.Routes()
	}
}
