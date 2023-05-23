package sqlerr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/slice"
)

// Error is a helper to convert SQL errors into friendly errors. It returns nil if no error is
// found.
// TODO: refactor this to use maps instead of if-else statements.
func Error[Model fmt.Stringer](err error, model Model) error {
	if err == nil {
		return nil
	}

	var errs *errutil.Error

	if strings.Contains(err.Error(), notFoundPattern) {
		return errutil.Append(errs, newErrNotFound(model.String()))
	}

	if strings.Contains(err.Error(), uniqueConstraintPattern) {
		return errutil.Append(errs, uniqueConstraint(err))
	}

	if strings.Contains(err.Error(), foreignKeyConstraintPattern) {
		return errutil.Append(errs, foreignKeyConstraint(err))
	}

	if strings.Contains(err.Error(), notNullConstraintPattern) {
		return errutil.Append(errs, notNullConstraint(err))
	}

	if strings.Contains(err.Error(), checkConstraintPattern) {
		return errutil.Append(errs, checkConstraint(err))
	}

	if strings.Contains(err.Error(), contextDeadlinePattern) {
		return errutil.Append(errs, newErrTimeout())
	}

	if strings.Contains(err.Error(), tooManyClientsPattern) {
		panic(errutil.Append(errs, newErrUnavailable(model.String())))
	}

	if strings.Contains(err.Error(), noRowsAffectedPattern) {
		return errutil.Append(errs, newErrNoChanges())
	}

	if strings.Contains(err.Error(), failedToConnectPattern) {
		panic(errutil.Append(errs, newErrUnavailable("service")))
	}

	if strings.Contains(err.Error(), columnDoesNotExistPattern) {
		return errutil.Append(errs, newErrColumnNotFound(getColumnName(err)))
	}

	if strings.Contains(err.Error(), inputSyntaxPattern) {
		return errutil.Append(errs, newErrMalformedQuery())
	}

	if strings.Contains(err.Error(), missingWhereConditionsPattern) {
		return errutil.Append(errs, newErrMalformedQuery())
	}

	if strings.Contains(err.Error(), tooLongValuePattern) {
		reg := regexp.MustCompile(`varying\(([0-9]+)\)`)
		matches := reg.FindStringSubmatch(err.Error())

		if len(matches) == 0 {
			return errutil.Append(errs, newErrCheckConstraint("field", "too long"))
		}

		value := slice.
			Map(matches[1:], func(value string) int {
				v, err := strconv.Atoi(value)
				if err != nil {
					panic(
						fmt.Errorf(fmt.Sprintf("failed to convert '%s' to int", value)),
					)
				}

				return v
			}).
			Build()

		return errutil.Append(errs, newErrTooLong(getColumnName(err), value[0]))
	}

	return errutil.Append(errs, err)
}
