package validation

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/validator"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	v "github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

// validate is the instance of the Validate.
var validate = v.New()

// validatorImpl implements the Validator interface.
type validatorImpl struct{}

// newValidator creates a newValidator Validator.
func newValidator() validator.Validator {
	return &validatorImpl{}
}

// Validate implements the Validator interface.
func (va validatorImpl) Validate(anStruct any) error {
	var errs *errutil.Error

	err := validate.Struct(anStruct)
	if err != nil {
		if verr, ok := err.(v.ValidationErrors); ok {
			for _, e := range verr {
				err, ok := errorMsgs[e.ActualTag()]
				if !ok {
					err = defaultError
				}

				errs = errutil.Append(errs, err(strcase.ToLowerCamel(e.Field()), e.Param()))
			}
		} else {
			return nil
		}
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
