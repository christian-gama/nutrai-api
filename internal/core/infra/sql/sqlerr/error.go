package sqlerr

import (
	"fmt"
	"strings"
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
		return newErrColumnNotFound(getColumnName(err))
	},

	contextDeadlinePattern: func(err error, model fmt.Stringer) error {
		return newErrTimeout()
	},

	failedToConnectPattern: func(err error, model fmt.Stringer) error {
		panic(newErrUnavailable("service"))
	},

	foreignKeyConstraintPattern: func(err error, model fmt.Stringer) error {
		return foreignKeyConstraint(err)
	},

	inputSyntaxPattern: func(err error, model fmt.Stringer) error {
		return newErrMalformedQuery()
	},

	missingWhereConditionsPattern: func(err error, model fmt.Stringer) error {
		return newErrMalformedQuery()
	},

	noRowsAffectedPattern: func(err error, model fmt.Stringer) error {
		return newErrNoChanges()
	},

	notFoundPattern: func(err error, model fmt.Stringer) error {
		return newErrNotFound(model.String())
	},

	notNullConstraintPattern: func(err error, model fmt.Stringer) error {
		return notNullConstraint(err)
	},

	tooLongValuePattern: func(err error, model fmt.Stringer) error {
		return tooLongConstraint(err)
	},

	tooManyClientsPattern: func(err error, model fmt.Stringer) error {
		panic(newErrUnavailable("service"))
	},

	uniqueConstraintPattern: func(err error, model fmt.Stringer) error {
		return uniqueConstraint(err)
	},
}
