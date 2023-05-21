package suite

import (
	"testing"

	"github.com/christian-gama/nutrai-api/testutils/sqlutil"
	"github.com/christian-gama/nutrai-api/testutils/suite/asserts"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Suite is the base suite for all test suites. It provides helper methods for
// testing.
type Suite struct {
	suite.Suite
}

// Skip skips the test.
func (s *Suite) Skip(name string, f func()) bool {
	return s.Run(name, func() {
		s.T().Skip()
	})
}

// Todo marks the test as TODO, skipping it and printing a message.
func (s *Suite) Todo(name string, f func()) bool {
	return s.Run(name, func() {
		s.T().Skipf("TODO: %s", name)
	})
}

// ErrorAsNotFound checks if the error is an ErrNotFound.
func (s *Suite) ErrorAsNotFound(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsNotFound(s.T(), err, msgAndArgs...)
}

// ErrorAsInvalid checks if the error is an ErrInvalid.
func (s *Suite) ErrorAsInvalid(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsInvalid(s.T(), err, msgAndArgs...)
}

// ErrorAsRequired checks if the error is an ErrRequired.
func (s *Suite) ErrorAsRequired(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsRequired(s.T(), err, msgAndArgs...)
}

// ErrorAsInternal checks if the error is an ErrInternal.
func (s *Suite) ErrorAsInternal(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsInternal(s.T(), err, msgAndArgs...)
}

// HasNotChanged checks if the given objects are equal.
func (s *Suite) HasNotChanged(
	oldObj, newObj any,
	msgAndArgs ...any,
) bool {
	return asserts.HasNotChanged(s.T(), oldObj, newObj, msgAndArgs...)
}

// HasChanged checks if the given object has changed.
func (s *Suite) HasChanged(
	oldObj, newObj any,
	msgAndArgs ...any,
) bool {
	return asserts.HasChanged(s.T(), oldObj, newObj, msgAndArgs...)
}

// SuiteWithSQLConn is the base suite for all test suites that need a SQL connection.
// It provides helper methods for testing and wraps each test in a transaction, rolling
// back the transaction after the test is done.
type SuiteWithSQLConn struct {
	Suite
}

// Run runs a test in a transaction and rolls back the transaction after the test is done.
// It is a wrapper around suite.Suite.Run.
func (s *SuiteWithSQLConn) Run(name string, f func(tx *gorm.DB)) bool {
	return s.Suite.Run(name, func() {
		sqlutil.Transaction(s.Fail, func(tx *gorm.DB) {
			f(tx)
		})
	})
}

// Skip skips the test.
func (s *SuiteWithSQLConn) SQLCount(
	db *gorm.DB,
	schema schema.Tabler,
	expectedCount int,
	msgAndArgs ...any,
) bool {
	return asserts.SQLCount(s.T(), db, schema, expectedCount, msgAndArgs...)
}

// SQLRecordExist checks if the given schema exists in the database.
func (s *SuiteWithSQLConn) SQLRecordExist(
	db *gorm.DB,
	schema schema.Tabler,
	msgAndArgs ...any,
) bool {
	return asserts.SQLRecordExist(s.T(), db, schema, msgAndArgs...)
}

// SQLRecordDoesNotExist checks if the given schema does not exist in the database.
func (s *SuiteWithSQLConn) SQLRecordDoesNotExist(
	db *gorm.DB,
	schema schema.Tabler,
	msgAndArgs ...any,
) bool {
	return asserts.SQLRecordDoesNotExist(s.T(), db, schema, msgAndArgs...)
}

func TestSetupTestsSuite(t *testing.T) {
	t.Helper()
	suite.Run(t, new(SuiteWithSQLConn))
}
