package sqlerr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	"github.com/iancoleman/strcase"
)

// checkConstraint is an error that occurs when a check constraint is violated.
// Example: "new row for relation \"user\" violates check constraint \"chk__email__email\"".
func checkConstraint(err error) error {
	msg := err.Error()

	// chk__column__message
	chkRegexp := regexp.MustCompile(`chk__(\w+)__(\w+)`)

	if !chkRegexp.MatchString(msg) {
		return err
	}

	matches := chkRegexp.FindStringSubmatch(msg)

	columnName := matches[1]
	message := matches[2]

	return errors.Invalid(
		strcase.ToLowerCamel(columnName),
		strings.ReplaceAll(message, "_", " "),
	)
}

// foreignKeyConstraint is an error that occurs when a foreign key constraint is violated.
// Example: "insert or update on table \"user\" violates foreign key constraint
// \"fk__role_id__roles.id\"".
func foreignKeyConstraint(err error) error {
	msg := err.Error()

	// \"table\" violates foreign key constraint \"fk__column__referenced_table.referenced_column\""
	/* \"promotion_keywords\" violates foreign key constraint \"fk__promotion_id__promotions.id\" */
	fkRegexp := regexp.MustCompile(
		`"(\w+)" violates foreign key constraint "fk__(\w+)__(\w+)\.(\w+)"`,
	)

	if !fkRegexp.MatchString(msg) {
		return err
	}

	matches := fkRegexp.FindStringSubmatch(msg)

	table := getFriendlyTableName(matches[1])
	column := matches[2]
	namespace := fmt.Sprintf("%s.%s", strcase.ToLowerCamel(table), strcase.ToLowerCamel(column))

	return errors.NotFound(namespace)
}

// notNullConstraint is an error that occurs when a not-null constraint is violated.
// Example: "pq: null value in column \"email\" violates not-null constraint".
func notNullConstraint(err error) error {
	str := err.Error()

	// value in column "column_name" violates not-null constraint
	nnRegexp := regexp.MustCompile(`value in column "(\w+)" violates not-null constraint`)

	if !nnRegexp.MatchString(str) {
		return notNullConstraintOfRelation(err)
	}

	matches := nnRegexp.FindStringSubmatch(str)

	columnName := matches[1]

	return errors.Required(columnName)
}

// notNullConstraintOfRelation is an error that occurs when a not-null constraint is violated in a
// relation.
func notNullConstraintOfRelation(err error) error {
	str := err.Error()

	// null value in column "column_name" of relation "relation_name" violates not-null constraint
	nnRegexp := regexp.MustCompile(
		`null value in column "(\w+)" of relation "(\w+)" violates not-null constraint`,
	)

	if !nnRegexp.MatchString(str) {
		return err
	}

	matches := nnRegexp.FindStringSubmatch(str)

	columnName := strcase.ToLowerCamel(matches[1])
	relationName := strcase.ToLowerCamel(matches[2])
	fullName := fmt.Sprintf("%s.%s", relationName, columnName)

	return errors.Required(fullName)
}

// tooLongConstraint is an error that occurs when a field is too long.
func tooLongConstraint(err error) error {
	reg := regexp.MustCompile(`varying\(([0-9]+)\)`)
	matches := reg.FindStringSubmatch(err.Error())

	if len(matches) == 0 {
		return errors.Invalid("", "too long")
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

	return errors.Invalid(getColumnName(err), fmt.Sprintf("too long (max %d)", value))
}

// uniqueConstraint is an error that occurs when a unique constraint is violated.
// Example: "pq: duplicate key value violates unique constraint \"uidx__email__users\"".
func uniqueConstraint(err error) error {
	msg := err.Error()

	// uidx__table__column
	uidxRegexp := regexp.MustCompile(`uidx__(\w+)__(\w+)`)

	// If the error message does not match the unique constraint regexp from uidx, then it may
	// be a unique constraint from a check constraint.
	if !uidxRegexp.MatchString(msg) {
		err := checkConstraint(err)
		if err != nil {
			chkRegexp := regexp.MustCompile(`chk__(\w+)__(\w+)`)

			if !chkRegexp.MatchString(msg) {
				return err
			}

			matches := chkRegexp.FindStringSubmatch(msg)

			columnName := matches[0]

			return errors.AlreadyExists(columnName)
		}
	}

	matches := uidxRegexp.FindStringSubmatch(msg)

	columnName := matches[2]

	return errors.AlreadyExists(columnName)
}
