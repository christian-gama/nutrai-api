package sql

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

func ForeignKeyConstraint(err error) error {
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

	table := matches[1]
	column := matches[2]
	namespace := fmt.Sprintf("%s.%s", strcase.ToLowerCamel(table), strcase.ToLowerCamel(column))
	referencedTable := matches[3]

	return NewErrForeignKeyConstraint(namespace, referencedTable)
}

func CheckConstraint(err error) error {
	msg := err.Error()

	// chk__column__message
	chkRegexp := regexp.MustCompile(`chk__(\w+)__(\w+)`)

	if !chkRegexp.MatchString(msg) {
		return err
	}

	matches := chkRegexp.FindStringSubmatch(msg)

	columnName := matches[1]
	message := matches[2]

	return NewErrCheckConstraint(
		strcase.ToLowerCamel(columnName),
		strings.ReplaceAll(message, "_", " "),
	)
}

func UniqueConstraint(err error) error {
	msg := err.Error()

	// uidx__table__column
	uidxRegexp := regexp.MustCompile(`uidx__(\w+)__(\w+)`)

	if !uidxRegexp.MatchString(msg) {
		return CheckConstraint(err)
	}

	matches := uidxRegexp.FindStringSubmatch(msg)

	columnName := matches[2]

	return NewErrUniqueConstraint(strcase.ToLowerCamel(columnName))
}

func NotNullConstraint(err error) error {
	str := err.Error()

	// value in column "column_name" violates not-null constraint
	nnRegexp := regexp.MustCompile(`value in column "(\w+)" violates not-null constraint`)

	if !nnRegexp.MatchString(str) {
		return NotNullConstraintOfRelation(err)
	}

	matches := nnRegexp.FindStringSubmatch(str)

	columnName := matches[1]

	return NewErrNotNullConstraint(strcase.ToLowerCamel(columnName))
}

func NotNullConstraintOfRelation(err error) error {
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
	output := fmt.Sprintf("%s.%s", relationName, columnName)

	return NewErrNotNullConstraint(output)
}
