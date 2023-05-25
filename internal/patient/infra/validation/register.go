package validation

import "github.com/christian-gama/nutrai-api/internal/core/infra/validation"

// Register is the function that registers the validation rules for this module.
func Register() {
	validation.RegisterAlias("patient_weight_kg", "number,min=30,max=600")
	validation.RegisterAlias("patient_height_m", "number,min=1,max=3")
	validation.RegisterAlias("patient_age", "number,min=18,max=100")
	validation.RegisterAlias("patient_allergy_name", "max=100")
}
