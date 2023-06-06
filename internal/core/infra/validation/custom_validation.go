package validation

import (
	"reflect"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/core/infra/log"

	"github.com/christian-gama/nutrai-api/internal/core/infra/validation/validators"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/go-playground/validator/v10"
)

func validateFilter(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	params := strings.Split(strings.TrimSpace(fl.Param()), " ")
	if len(params) == 0 || params[0] == "" {
		log.Panic(errors.InternalServerError("the filter tag must have parameters"))
	}

	return validators.Filter(fl.Field().String(), params)
}

func validateSort(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	params := strings.Split(fl.Param(), " ")
	if len(params) == 0 || params[0] == "" {
		log.Panic(errors.InternalServerError("the sort tag must have parameters"))
	}

	return validators.Sort(fl.Field().String(), params)
}

func validatePreload(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	params := strings.Split(strings.TrimSpace(fl.Param()), " ")
	if len(params) == 0 || params[0] == "" {
		log.Panic(errors.InternalServerError("the preload tag must have parameters"))
	}

	return validators.Preload(fl.Field().String(), params)
}
