package sqlerr

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

// errNotFound is the error returned when a resource is not found.
type errNotFound struct {
	Resource string
}

// newErrNotFound returns a new ErrNotFound error.
// Output: "could not find resource".
func newErrNotFound(resource string, args ...any) error {
	return &errNotFound{Resource: resource}
}

// Error implements error.
func (e *errNotFound) Error() string {
	return fmt.Sprintf("could not find %s", e.Resource)
}

// errUniqueConstraint is the error returned when a unique constraint is violated.
type errUniqueConstraint struct {
	Field string
}

// newErrUniqueConstraint returns a new ErrUniqueConstraint error.
// Output: "field already exists".
func newErrUniqueConstraint(field string) error {
	return &errUniqueConstraint{Field: field}
}

// Error implements error.
func (e *errUniqueConstraint) Error() string {
	return fmt.Sprintf("%s already exists", e.Field)
}

// errForeignKeyConstraint is the error returned when a foreign key constraint is violated.
type errForeignKeyConstraint struct {
	Field          string
	ReferenceTable string
}

// newErrForeignKeyConstraint returns a new ErrForeignKeyConstraint error.
// Output: "field does not exist".
func newErrForeignKeyConstraint(field string, referenceTable string) error {
	return &errForeignKeyConstraint{Field: field, ReferenceTable: referenceTable}
}

// Error implements error.
func (e *errForeignKeyConstraint) Error() string {
	return fmt.Sprintf("%s does not exist on %s", e.Field, strcase.ToLowerCamel(e.ReferenceTable))
}

// errCheckConstraint is the error returned when a check constraint is violated.
type errCheckConstraint struct {
	Field  string
	Reason string
	args   []any
}

// newErrCheckConstraint returns a new ErrCheckConstraint error.
// Output: "field is invalid: reason".
func newErrCheckConstraint(field string, reason string, args ...any) error {
	return &errCheckConstraint{Field: field, Reason: reason, args: args}
}

// Error implements error.
func (e *errCheckConstraint) Error() string {
	reasonWithArgs := fmt.Sprintf(e.Reason, e.args...)
	return fmt.Sprintf("%s is invalid: %s", e.Field, reasonWithArgs)
}

// errNotNullConstraint is the error returned when a not null constraint is violated.
type errNotNullConstraint struct {
	Field string
}

// newErrNotNullConstraint returns a new ErrNotNullConstraint error.
// Output: "field cannot be null".
func newErrNotNullConstraint(field string) error {
	return &errNotNullConstraint{Field: field}
}

// Error implements error.
func (e *errNotNullConstraint) Error() string {
	return fmt.Sprintf("%s cannot be null", e.Field)
}

// errTimeout is the error returned when a timeout occurs.
type errTimeout struct{}

// newErrTimeout returns a new ErrTimeout error.
// Output: "timeout occurred".
func newErrTimeout() error {
	return &errTimeout{}
}

// Error implements error.
func (e *errTimeout) Error() string {
	return "timeout occurred"
}

// errUnavailable is the error returned when a resource is unavailable.
type errUnavailable struct {
	Resource string
}

// newErrUnavailable returns a new ErrUnavailable error.
// Output: "resource is unavailable".
func newErrUnavailable(resource string) error {
	return &errUnavailable{Resource: resource}
}

// Error implements error.
func (e *errUnavailable) Error() string {
	return fmt.Sprintf("%s is unavailable", e.Resource)
}

// errNoChanges is the error returned when no changes were made.
type errNoChanges struct{}

// newErrNoChanges returns a new ErrNoChanges error.
// Output: "no changes were made".
func newErrNoChanges() error {
	return &errNoChanges{}
}

// Error implements error.
func (e *errNoChanges) Error() string {
	return "no changes were made"
}

// errTooLong is the error returned when a string is too long.
type errTooLong struct {
	Field string
	Max   int
}

// newErrTooLong returns a new ErrTooLong error.
// Output: "field is too long (max: max)".
func newErrTooLong(field string, max int) error {
	return &errTooLong{Field: field, Max: max}
}

// Error implements error.
func (e *errTooLong) Error() string {
	return fmt.Sprintf("%s is too long (max: %d)", e.Field, e.Max)
}

// errMalformedQuery is the error returned when a query is malformed.
type errMalformedQuery struct{}

// newErrMalformedQuery returns a new ErrMalformedQuery error.
// Output: "could not proceed because of malformed query".
func newErrMalformedQuery() error {
	return &errMalformedQuery{}
}

// Error implements error.
func (e *errMalformedQuery) Error() string {
	return "could not proceed because of malformed query"
}

// errColumnNotFound is the error returned when a column is not found.
type errColumnNotFound struct {
	Field string
}

// newErrColumnNotFound returns a new ErrColumnNotFound error.
// Output: "could not find field column".
func newErrColumnNotFound(field string) error {
	return &errColumnNotFound{Field: field}
}

// Error implements error.
func (e *errColumnNotFound) Error() string {
	return fmt.Sprintf("could not find %s column", e.Field)
}
