package sql

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/slice"
)

var (
	errNotFound               = "record not found"
	errUniqueConstraint       = "violates unique constraint"
	errForeignKeyConstraint   = "violates foreign key constraint"
	errNotNullConstraint      = "violates not-null constraint"
	errCheckConstraint        = "violates check constraint"
	errContextDeadline        = "context deadline exceeded"
	errTooManyClients         = "sorry, too many clients already"
	errNoRowsAffected         = "no rows affected"
	errFailedToConnect        = "failed to connect to"
	errColumnDoesNotExist     = "SQLSTATE 42703"
	errInputSyntax            = "SQLSTATE 22P02"
	errMissingWhereConditions = "WHERE conditions required"
	errTooLongValue           = "SQLSTATE 22001"
)

// Error is a helper to convert SQL errors into friendly errors. It returns nil if no error is
// found.
// TODO: refactor this to use maps instead of if-else statements.
func Error[Model fmt.Stringer](err error, model Model) error {
	if err == nil {
		return nil
	}

	var errs *errutil.Error

	if strings.Contains(err.Error(), errNotFound) {
		return errutil.Append(errs, NewErrNotFound(model.String()))
	}

	if strings.Contains(err.Error(), errUniqueConstraint) {
		return errutil.Append(errs, UniqueConstraint(err))
	}

	if strings.Contains(err.Error(), errForeignKeyConstraint) {
		return errutil.Append(errs, ForeignKeyConstraint(err))
	}

	if strings.Contains(err.Error(), errNotNullConstraint) {
		return errutil.Append(errs, NotNullConstraint(err))
	}

	if strings.Contains(err.Error(), errCheckConstraint) {
		return errutil.Append(errs, CheckConstraint(err))
	}

	if strings.Contains(err.Error(), errContextDeadline) {
		return errutil.Append(errs, NewErrTimeout())
	}

	if strings.Contains(err.Error(), errTooManyClients) {
		panic(errutil.Append(errs, NewErrUnavailable(model.String())))
	}

	if strings.Contains(err.Error(), errNoRowsAffected) {
		return errutil.Append(errs, NewErrNoChanges())
	}

	if strings.Contains(err.Error(), errFailedToConnect) {
		panic(errutil.Append(errs, NewErrUnavailable("service")))
	}

	if strings.Contains(err.Error(), errColumnDoesNotExist) {
		return errutil.Append(errs, NewErrColumnNotFound(getColumnName(err)))
	}

	if strings.Contains(err.Error(), errInputSyntax) {
		return errutil.Append(errs, NewErrMalformedQuery())
	}

	if strings.Contains(err.Error(), errMissingWhereConditions) {
		return errutil.Append(errs, NewErrMalformedQuery())
	}

	if strings.Contains(err.Error(), errTooLongValue) {
		reg := regexp.MustCompile(`varying\(([0-9]+)\)`)
		matches := reg.FindStringSubmatch(err.Error())

		if len(matches) == 0 {
			return errutil.Append(errs, NewErrCheckConstraint("field", "too long"))
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

		return errutil.Append(errs, NewErrTooLong(getColumnName(err), value[0]))
	}

	return errutil.Append(errs, err)
}
