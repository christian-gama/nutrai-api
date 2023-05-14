package validation

import (
	"errors"
	"reflect"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/shared/infra/validation/validators"
	"github.com/go-playground/validator/v10"
)

func validateFilter(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	params := strings.Split(fl.Param(), " ")
	if len(params) == 0 {
		panic(errors.New("the filter tag must have parameters"))
	}

	return validators.Filter(fl.Field().String(), params)
}

func validateSort(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	params := strings.Split(fl.Param(), " ")
	if len(params) == 0 {
		panic(errors.New("the sort tag must have parameters"))
	}

	return validators.Sort(fl.Field().String(), params)
}

func validatePreload(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	params := strings.Split(fl.Param(), " ")
	if len(params) == 0 {
		panic(errors.New("the preload tag must have parameters"))
	}

	return validators.Preload(fl.Field().String(), params)
}
