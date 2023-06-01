package sqlerr

import (
	"fmt"
	"strings"

	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// Error is a helper to convert SQL errors into friendly errors. It returns nil if no error is
// found.
func Error[Model fmt.Stringer](err error, model Model) error {
	if err == nil {
		return nil
	}

	for pattern, fn := range errsMap {
		if !strings.Contains(err.Error(), pattern) {
			continue
		}

		return fn(err, model)
	}

	return err
}

var errsMap = map[string]func(err error, model fmt.Stringer) error{
	checkConstraintPattern: func(err error, model fmt.Stringer) error {
		return checkConstraint(err)
	},

	columnDoesNotExistPattern: func(err error, model fmt.Stringer) error {
		return errors.InternalServerError("column does not exist")
	},

	contextDeadlinePattern: func(err error, model fmt.Stringer) error {
		return errors.Timeout()
	},

	failedToConnectPattern: func(err error, model fmt.Stringer) error {
		return errors.Unavailable("service")
	},

	foreignKeyConstraintPattern: func(err error, model fmt.Stringer) error {
		return foreignKeyConstraint(err)
	},

	inputSyntaxPattern: func(err error, model fmt.Stringer) error {
		return errors.InternalServerError("invalid syntax")
	},

	missingWhereConditionsPattern: func(err error, model fmt.Stringer) error {
		return errors.InternalServerError("missing where conditions")
	},

	noRowsAffectedPattern: func(err error, model fmt.Stringer) error {
		return errors.NoChanges()
	},

	notFoundPattern: func(err error, model fmt.Stringer) error {
		return errors.NotFound(model.String())
	},

	notNullConstraintPattern: func(err error, model fmt.Stringer) error {
		return notNullConstraint(err)
	},

	tooLongValuePattern: func(err error, model fmt.Stringer) error {
		return tooLongConstraint(err)
	},

	tooManyClientsPattern: func(err error, model fmt.Stringer) error {
		return errors.Unavailable("service")
	},

	uniqueConstraintPattern: func(err error, model fmt.Stringer) error {
		return uniqueConstraint(err)
	},

	relationDoesNotExistPattern: func(err error, model fmt.Stringer) error {
		return errors.InternalServerError(
			"relation for %s does not exist - did you run migrations?",
			model.String(),
		)
	},
}
