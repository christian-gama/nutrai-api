package errutil

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/christian-gama/nutrai-api/pkg/reflection"
)

func MustBeNotEmpty(name string, value any) {
	if reflection.IsZero(value) {
		panic(errors.InternalServerError("%s must not be empty", name))
	}
}
