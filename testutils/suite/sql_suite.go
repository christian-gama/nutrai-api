package suite

import (
	"github.com/christian-gama/nutrai-api/testutils/sqlutil"
	"github.com/christian-gama/nutrai-api/testutils/suite/asserts"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

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
