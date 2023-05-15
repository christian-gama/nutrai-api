package sql_test

import (
	"errors"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/sql"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type ErrorSuite struct {
	suite.Suite
}

func TestErrorSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ErrorSuite))
}

func (s *ErrorSuite) TestError() {
	s.Run("Error returns nil if the error is nil", func() {
		err := sql.Error(nil, "resource")

		s.Nil(err)
	})

	s.Run("returns an ErrNotFound if the error is a record not found", func() {
		err := errors.New("record not found")

		err = sql.Error(err, "resource")

		var e *sql.ErrNotFound
		s.ErrorAs(err, &e)
	})

	s.Run("returns an ErrUniqueConstraint if the error is a unique constraint violation",
		func() {
			err := errors.New("violates unique constraint uidx__table__column")

			err = sql.Error(err, "resource")

			var e *sql.ErrUniqueConstraint
			s.ErrorAs(err, &e)
		},
	)

	s.Run(
		"returns an ErrForeignKeyConstraint if the error is a foreign key constraint violation",
		func() {
			err := errors.New(
				`"table" violates foreign key constraint "fk__column__refTable.refColumn"`,
			)

			err = sql.Error(err, "resource")

			var e *sql.ErrForeignKeyConstraint
			s.ErrorAs(err, &e)
		},
	)

	s.Run("returns an ErrNotNullConstraint if the error is a not-null constraint violation", func() {
		err := errors.New(`value in column "column" violates not-null constraint`)

		err = sql.Error(err, "resource")

		var e *sql.ErrNotNullConstraint
		s.ErrorAs(err, &e)
	},
	)

	s.Run("returns an ErrCheckConstraint if the error is a check constraint violation", func() {
		err := errors.New("violates check constraint chk__column__message")

		err = sql.Error(err, "resource")

		var e *sql.ErrCheckConstraint
		s.ErrorAs(err, &e)
	},
	)

	s.Run("returns an ErrTimeout if the error is a context deadline exceeded", func() {
		err := errors.New("context deadline exceeded")

		err = sql.Error(err, "resource")

		var e *sql.ErrTimeout
		s.ErrorAs(err, &e)
	})

	s.Run("returns an ErrUnavailable if the error is a connection refused", func() {
		s.Panics(func() {
			err := errors.New("sorry, too many clients already")

			err = sql.Error(err, "resource")

			var e *sql.ErrUnavailable
			s.ErrorAs(err, &e)
		})
	})

	s.Run("returns an ErrNoChanges if the error is a no rows affected", func() {
		err := errors.New("no rows affected")

		err = sql.Error(err, "resource")

		var e *sql.ErrNoChanges
		s.ErrorAs(err, &e)
	})

	s.Run("returns an ErrUnavailable if the error is a connection refused", func() {
		s.Panics(func() {
			err := errors.New("failed to connect to")

			err = sql.Error(err, "resource")

			var e *sql.ErrUnavailable
			s.ErrorAs(err, &e)
		})
	})

	s.Run("returns an ErrColumnNotFound if the error is a column does not exist", func() {
		err := errors.New("column \"id\" of relation \"resource\" does not exist SQLSTATE 42703")

		err = sql.Error(err, "resource")

		var e *sql.ErrColumnNotFound
		s.ErrorAs(err, &e)
	})

	s.Run(
		"returns an ErrColumnNotFound with generic field name if the error is can't extract params",
		func() {
			err := errors.New("SQLSTATE 42703")

			err = sql.Error(err, "resource")

			var e *sql.ErrColumnNotFound
			s.ErrorAs(err, &e)
		},
	)

	s.Run("returns an ErrMalformedQuery if the error is a input syntax", func() {
		err := errors.New("SQLSTATE 22P02")

		err = sql.Error(err, "resource")

		var e *sql.ErrMalformedQuery
		s.ErrorAs(err, &e)
	})

	s.Run("Error returns an ErrMalformedQuery if the error is a missing where conditions", func() {
		err := errors.New("WHERE conditions required")

		err = sql.Error(err, "resource")

		var e *sql.ErrMalformedQuery
		s.ErrorAs(err, &e)
	})

	s.Run("returns an ErrTooLong if the error is a too long value", func() {
		err := errors.New("value too long for type character varying(255) SQLSTATE 22001")

		err = sql.Error(err, "resource")

		var e *sql.ErrTooLong
		s.ErrorAs(err, &e)
	})

	s.Run("returns an ErrTooLong if the error is a too long value and cannot extract param", func() {
		err := errors.New("SQLSTATE 22001")

		err = sql.Error(err, "resource")

		var e *sql.ErrCheckConstraint
		s.ErrorAs(err, &e)
	},
	)

	s.Run("returns the original error if it is a unknown error", func() {
		err := errors.New("any other error")

		err = sql.Error(err, "resource")

		s.Contains(err.Error(), "any other error")
	})
}
