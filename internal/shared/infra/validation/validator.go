package validation

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	v "github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

type Validator interface {
	// Validate validates the given struct.
	Validate(anStruct any) error
}

// validate is the instance of the Validate.
var validate = v.New()

// validatorImpl implements the Validator interface.
type validatorImpl struct{}

// New creates a new Validator.
func New() Validator {
	validate.RegisterValidation("sort", validateSort)
	validate.RegisterValidation("filter", validateFilter)
	validate.RegisterValidation("preload", validatePreload)
	validate.RegisterAlias("query", "dive,omitempty")

	return &validatorImpl{}
}

// Validate implements the Validator interface.
func (va validatorImpl) Validate(anStruct any) error {
	var errs *errutil.Error

	err := validate.Struct(anStruct)
	if err != nil {
		if verr, ok := err.(v.ValidationErrors); ok {
			for _, e := range verr {
				err, ok := ErrorMap[e.ActualTag()]
				if !ok {
					err = DefaultError
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

func RegisterAlias(tag, alias string) {
	validate.RegisterAlias(tag, alias)
}
