package auth

import (
	"github.com/christian-gama/nutrai-api/internal/auth/api/http"
	. "github.com/christian-gama/nutrai-api/internal/auth/module"
	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
)

// Init is the function that initializes this module.
func Init() (*module.Module, func()) {
	return Module, func() {
		// Add the logic to initialize this module here:
		http.Routes()
	}
}
