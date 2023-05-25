package validation

import "github.com/christian-gama/nutrai-api/internal/core/infra/validation"

// Register is the function that registers the validation rules for this module.
func Register() {
	validation.RegisterAlias("user_password", "min=8,max=32")
	validation.RegisterAlias("user_name", "min=2,max=100")
}
