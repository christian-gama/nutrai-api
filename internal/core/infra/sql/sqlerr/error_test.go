package sqlerr

import (
	"errors"
	"testing"

	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type ErrorSuite struct {
	suite.Suite
}

func TestErrorSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ErrorSuite))
}

type testModel struct{}

func (t *testModel) String() string {
	return "resource"
}

func (t *testModel) Validate() (*testModel, error) {
	return nil, nil
}

var TestModel = &testModel{}

func (s *ErrorSuite) TestError() {
	s.Run("Error returns nil if the error is nil", func() {
		err := Error(nil, TestModel)

		s.Nil(err)
	})

	s.Run("returns an ErrNotFound if the error is a record not found", func() {
		err := errors.New("record not found")

		err = Error(err, TestModel)

		var e *errNotFound
		s.ErrorAs(err, &e)
	})

	s.Run("returns an ErrUniqueConstraint if the error is a unique constraint violation",
		func() {
			err := errors.New("violates unique constraint uidx__table__column")

			err = Error(err, TestModel)

			var e *errUniqueConstraint
			s.ErrorAs(err, &e)
		},
	)

	s.Run(
		"returns an ErrForeignKeyConstraint if the error is a foreign key constraint violation",
		func() {
			err := errors.New(
				`"table" violates foreign key constraint "fk__column__refTable.refColumn"`,
			)

			err = Error(err, TestModel)

			var e *errForeignKeyConstraint
			s.ErrorAs(err, &e)
		},
	)

	s.Run(
		"returns an ErrNotNullConstraint if the error is a not-null constraint violation",
		func() {
			err := errors.New(`value in column "column" violates not-null constraint`)

			err = Error(err, TestModel)

			var e *errNotNullConstraint
			s.ErrorAs(err, &e)
		},
	)

	s.Run("returns an ErrCheckConstraint if the error is a check constraint violation", func() {
		err := errors.New("violates check constraint chk__column__message")

		err = Error(err, TestModel)

		var e *errCheckConstraint
		s.ErrorAs(err, &e)
	},
	)

	s.Run("returns an ErrTimeout if the error is a context deadline exceeded", func() {
		err := errors.New("context deadline exceeded")

		err = Error(err, TestModel)

		var e *errTimeout
		s.ErrorAs(err, &e)
	})

	s.Run("returns an ErrUnavailable if the error is a connection refused", func() {
		s.Panics(func() {
			err := errors.New("sorry, too many clients already")

			err = Error(err, TestModel)

			var e *errUnavailable
			s.ErrorAs(err, &e)
		})
	})

	s.Run("returns an ErrNoChanges if the error is a no rows affected", func() {
		err := errors.New("no rows affected")

		err = Error(err, TestModel)

		var e *errNoChanges
		s.ErrorAs(err, &e)
	})

	s.Run("returns an ErrUnavailable if the error is a connection refused", func() {
		s.Panics(func() {
			err := errors.New("failed to connect to")

			err = Error(err, TestModel)

			var e *errUnavailable
			s.ErrorAs(err, &e)
		})
	})

	s.Run("returns an ErrColumnNotFound if the error is a column does not exist", func() {
		err := errors.New("column \"id\" of relation \"resource\" does not exist SQLSTATE 42703")

		err = Error(err, TestModel)

		var e *errColumnNotFound
		s.ErrorAs(err, &e)
	})

	s.Run(
		"returns an ErrColumnNotFound with generic field name if the error is can't extract params",
		func() {
			err := errors.New("SQLSTATE 42703")

			err = Error(err, TestModel)

			var e *errColumnNotFound
			s.ErrorAs(err, &e)
		},
	)

	s.Run("returns an ErrMalformedQuery if the error is a input syntax", func() {
		err := errors.New("SQLSTATE 22P02")

		err = Error(err, TestModel)

		var e *errMalformedQuery
		s.ErrorAs(err, &e)
	})

	s.Run("Error returns an ErrMalformedQuery if the error is a missing where conditions", func() {
		err := errors.New("WHERE conditions required")

		err = Error(err, TestModel)

		var e *errMalformedQuery
		s.ErrorAs(err, &e)
	})

	s.Run("returns an ErrTooLong if the error is a too long value", func() {
		err := errors.New("value too long for type character varying(255) SQLSTATE 22001")

		err = Error(err, TestModel)

		var e *errTooLong
		s.ErrorAs(err, &e)
	})

	s.Run(
		"returns an ErrTooLong if the error is a too long value and cannot extract param",
		func() {
			err := errors.New("SQLSTATE 22001")

			err = Error(err, TestModel)

			var e *errCheckConstraint
			s.ErrorAs(err, &e)
		},
	)

	s.Run("returns the original error if it is a unknown error", func() {
		err := errors.New("any other error")

		err = Error(err, TestModel)

		s.Contains(err.Error(), "any other error")
	})
}
