package validation

import "github.com/christian-gama/nutrai-api/internal/core/domain/validator"

var validatorSingleton validator.Validator

func MakeValidator() validator.Validator {
	if validatorSingleton == nil {
		validatorSingleton = newValidator()
	}

	return validatorSingleton
}
