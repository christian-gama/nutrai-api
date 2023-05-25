package sqlerr

import (
	"fmt"
	"strings"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
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
		return errutil.Repository(checkConstraint(err).Error())
	},

	columnDoesNotExistPattern: func(err error, model fmt.Stringer) error {
		return errutil.Repository(newErrColumnNotFound(getColumnName(err)).Error())
	},

	contextDeadlinePattern: func(err error, model fmt.Stringer) error {
		panic(newErrTimeout())
	},

	failedToConnectPattern: func(err error, model fmt.Stringer) error {
		panic(newErrUnavailable("service"))
	},

	foreignKeyConstraintPattern: func(err error, model fmt.Stringer) error {
		return errutil.Repository(foreignKeyConstraint(err).Error())
	},

	inputSyntaxPattern: func(err error, model fmt.Stringer) error {
		panic(newErrMalformedQuery())
	},

	missingWhereConditionsPattern: func(err error, model fmt.Stringer) error {
		panic(newErrMalformedQuery())
	},

	noRowsAffectedPattern: func(err error, model fmt.Stringer) error {
		return errutil.Repository(newErrNoChanges().Error())
	},

	notFoundPattern: func(err error, model fmt.Stringer) error {
		return errutil.Repository(newErrNotFound(model.String()).Error())
	},

	notNullConstraintPattern: func(err error, model fmt.Stringer) error {
		return errutil.Repository(notNullConstraint(err).Error())
	},

	tooLongValuePattern: func(err error, model fmt.Stringer) error {
		return errutil.Repository(tooLongConstraint(err).Error())
	},

	tooManyClientsPattern: func(err error, model fmt.Stringer) error {
		panic(newErrUnavailable("service"))
	},

	uniqueConstraintPattern: func(err error, model fmt.Stringer) error {
		return errutil.Repository(uniqueConstraint(err).Error())
	},
}
