package errutil

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/reflection"
)

func MustBeNotEmpty(name string, value any) {
	if reflection.IsZero(value) {
		panic(fmt.Errorf("%s must not be empty", name))
	}
}
