package sql

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

// ErrNotFound is the error returned when a resource is not found.
type ErrNotFound struct {
	Resource string
}

// NewErrNotFound returns a new ErrNotFound error.
// Output: "could not find resource".
func NewErrNotFound(resource string, args ...any) error {
	return &ErrNotFound{Resource: resource}
}

// Error implements error.
func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("could not find %s", e.Resource)
}

// ErrUniqueConstraint is the error returned when a unique constraint is violated.
type ErrUniqueConstraint struct {
	Field string
}

// NewErrUniqueConstraint returns a new ErrUniqueConstraint error.
// Output: "field already exists".
func NewErrUniqueConstraint(field string) error {
	return &ErrUniqueConstraint{Field: field}
}

// Error implements error.
func (e *ErrUniqueConstraint) Error() string {
	return fmt.Sprintf("%s already exists", e.Field)
}

// ErrForeignKeyConstraint is the error returned when a foreign key constraint is violated.
type ErrForeignKeyConstraint struct {
	Field          string
	ReferenceTable string
}

// NewErrForeignKeyConstraint returns a new ErrForeignKeyConstraint error.
// Output: "field does not exist".
func NewErrForeignKeyConstraint(field string, referenceTable string) error {
	return &ErrForeignKeyConstraint{Field: field, ReferenceTable: referenceTable}
}

// Error implements error.
func (e *ErrForeignKeyConstraint) Error() string {
	return fmt.Sprintf("%s does not exist on %s", e.Field, strcase.ToLowerCamel(e.ReferenceTable))
}

// ErrCheckConstraint is the error returned when a check constraint is violated.
type ErrCheckConstraint struct {
	Field  string
	Reason string
	args   []any
}

// NewErrCheckConstraint returns a new ErrCheckConstraint error.
// Output: "field is invalid: reason".
func NewErrCheckConstraint(field string, reason string, args ...any) error {
	return &ErrCheckConstraint{Field: field, Reason: reason, args: args}
}

// Error implements error.
func (e *ErrCheckConstraint) Error() string {
	reasonWithArgs := fmt.Sprintf(e.Reason, e.args...)
	return fmt.Sprintf("%s is invalid: %s", e.Field, reasonWithArgs)
}

// ErrNotNullConstraint is the error returned when a not null constraint is violated.
type ErrNotNullConstraint struct {
	Field string
}

// NewErrNotNullConstraint returns a new ErrNotNullConstraint error.
// Output: "field cannot be null".
func NewErrNotNullConstraint(field string) error {
	return &ErrNotNullConstraint{Field: field}
}

// Error implements error.
func (e *ErrNotNullConstraint) Error() string {
	return fmt.Sprintf("%s cannot be null", e.Field)
}

// ErrTimeout is the error returned when a timeout occurs.
type ErrTimeout struct{}

// NewErrTimeout returns a new ErrTimeout error.
// Output: "timeout occurred".
func NewErrTimeout() error {
	return &ErrTimeout{}
}

// Error implements error.
func (e *ErrTimeout) Error() string {
	return "timeout occurred"
}

// ErrUnavailable is the error returned when a resource is unavailable.
type ErrUnavailable struct {
	Resource string
}

// NewErrUnavailable returns a new ErrUnavailable error.
// Output: "resource is unavailable".
func NewErrUnavailable(resource string) error {
	return &ErrUnavailable{Resource: resource}
}

// Error implements error.
func (e *ErrUnavailable) Error() string {
	return fmt.Sprintf("%s is unavailable", e.Resource)
}

// ErrNoChanges is the error returned when no changes were made.
type ErrNoChanges struct{}

// NewErrNoChanges returns a new ErrNoChanges error.
// Output: "no changes were made".
func NewErrNoChanges() error {
	return &ErrNoChanges{}
}

// Error implements error.
func (e *ErrNoChanges) Error() string {
	return "no changes were made"
}

// ErrTooLong is the error returned when a string is too long.
type ErrTooLong struct {
	Field string
	Max   int
}

// NewErrTooLong returns a new ErrTooLong error.
// Output: "field is too long (max: max)".
func NewErrTooLong(field string, max int) error {
	return &ErrTooLong{Field: field, Max: max}
}

// Error implements error.
func (e *ErrTooLong) Error() string {
	return fmt.Sprintf("%s is too long (max: %d)", e.Field, e.Max)
}

// ErrMalformedQuery is the error returned when a query is malformed.
type ErrMalformedQuery struct{}

// NewErrMalformedQuery returns a new ErrMalformedQuery error.
// Output: "could not proceed because of malformed query".
func NewErrMalformedQuery() error {
	return &ErrMalformedQuery{}
}

// Error implements error.
func (e *ErrMalformedQuery) Error() string {
	return "could not proceed because of malformed query"
}

// ErrColumnNotFound is the error returned when a column is not found.
type ErrColumnNotFound struct {
	Field string
}

// NewErrColumnNotFound returns a new ErrColumnNotFound error.
// Output: "could not find field column".
func NewErrColumnNotFound(field string) error {
	return &ErrColumnNotFound{Field: field}
}

// Error implements error.
func (e *ErrColumnNotFound) Error() string {
	return fmt.Sprintf("could not find %s column", e.Field)
}
