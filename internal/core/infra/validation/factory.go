package validation

var validatorSingleton Validator

func MakeValidator() Validator {
	if validatorSingleton == nil {
		validatorSingleton = New()
	}

	return validatorSingleton
}
