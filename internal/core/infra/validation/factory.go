package validation

import "github.com/christian-gama/nutrai-api/internal/core/domain/validator"

var defaultValidator validator.Validator

func MakeValidator() validator.Validator {
	if defaultValidator == nil {
		defaultValidator = newValidator()
	}

	return defaultValidator
}
