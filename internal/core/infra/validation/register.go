package validation

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	v "github.com/go-playground/validator/v10"
)

func init() {
	RegisterAlias("query", "dive,omitempty")

	RegisterValidation(
		"sort",
		validateSort,
		func(field string, param string) string {
			return fmt.Sprintf(
				"must have the format 'field:asc|desc' and only one of the fields are allowed: %s",
				param,
			)
		})

	RegisterValidation(
		"filter",
		validateFilter,
		func(field string, param string) string {
			return fmt.Sprintf(
				"must have the format 'field=name,op=%s,value=' and only one of the fields are allowed: %s",
				querying.AllowedFilterOperators(),
				param,
			)
		})

	RegisterValidation(
		"preload",
		validatePreload,
		func(field string, param string) string {
			return fmt.Sprintf(
				"must have one of the fields: %s",
				param,
			)
		})
}

// RegisterValidation registers a validation for a specified tag. It will also register an error
// message for that tag.
func RegisterValidation(
	tag string,
	fn v.Func,
	errMsgFunc func(field string, param string) string,
	callValidationEvenIfNull ...bool,
) error {
	if errMsgFunc == nil {
		errorMsgs[tag] = defaultError
	} else {
		errorMsgs[tag] = func(field, param string) error {
			return errutil.Invalid(field, errMsgFunc(field, param))
		}
	}

	return validate.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

// RegisterAlias registers an alias for a tag.
func RegisterAlias(alias, tag string) {
	validate.RegisterAlias(alias, tag)
}
