package sqlutil

import (
	gosql "database/sql"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
	"gorm.io/gorm"
)

// Transaction is a helper function to run a transaction in a test. It will
// rollback automatically in the end of the function.
func Transaction(
	failFn func(failureMessage string, msgAndArgs ...interface{}) bool,
	fn func(tx *gorm.DB),
) {
	db := conn.GetPsql()
	defer func() {
		conn.ClosePsql()
	}()

	tx := func(tx *gorm.DB) error {
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				failFn("test panicked", "panic: %v", r)
			}
		}()

		fn(tx)

		return errors.New("it will rollback automatically on error")
	}

	db.Transaction(tx, &gosql.TxOptions{Isolation: gosql.LevelSerializable})
}
