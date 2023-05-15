package sql_test

import (
	"errors"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/sql"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type ConstraintsSuite struct {
	suite.Suite
}

func TestConstraintsSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ConstraintsSuite))
}

func (s *ConstraintsSuite) TestForeignKeyConstraint() {
	s.Run("returns an ErrDoesNotExist if the error is a foreign key constraint violation", func() {
		err := errors.New(
			`"table" violates foreign key constraint "fk__column__refTable.refColumn"`,
		)

		err = sql.ForeignKeyConstraint(err)

		s.Error(err)
		s.Equal("table.column does not exist on refTable", err.Error())
	})

	s.Run("returns the original error if it is not a foreign key constraint violation", func() {
		err := errors.New("some error")

		err = sql.ForeignKeyConstraint(err)

		s.Equal("some error", err.Error())
	})
}

func (s *ConstraintsSuite) TestCheckConstraint() {
	s.Run("returns an ErrInvalid if the error is a check constraint violation", func() {
		err := errors.New("chk__column__message")

		err = sql.CheckConstraint(err)

		s.Error(err)
		s.Equal("column is invalid: message", err.Error())
	})

	s.Run("returns the original error if it is not a check constraint violation", func() {
		err := errors.New("some error")

		err = sql.CheckConstraint(err)

		s.Equal("some error", err.Error())
	})
}

func (s *ConstraintsSuite) TestUniqueConstraint() {
	s.Run("returns an ErrAlreadyExists if the error is a unique constraint violation", func() {
		err := errors.New("uidx__table__column")

		err = sql.UniqueConstraint(err)

		s.Error(err)
		s.Equal("column already exists", err.Error())
	})

	s.Run("returns the original error if it is not a unique constraint violation", func() {
		err := errors.New("some error")

		err = sql.UniqueConstraint(err)

		s.Equal("some error", err.Error())
	})

	s.Run(
		"returns an ErrIsInvalid if the error is a unique constraint violation but comes from chk",
		func() {
			err := errors.New("chk__column__message")

			err = sql.UniqueConstraint(err)

			s.Error(err)
			s.Equal("column is invalid: message", err.Error())
		},
	)
}

func (s *ConstraintsSuite) TestNotNullConstraint() {
	s.Run("returns an ErrRequired if the error is a not null constraint violation", func() {
		err := errors.New("null value in column \"column\" violates not-null constraint")

		err = sql.NotNullConstraint(err)

		s.Error(err)
		s.Equal("column cannot be null", err.Error())
	})

	s.Run("returns the original error if it is not a not null constraint violation", func() {
		err := errors.New("some error")

		err = sql.NotNullConstraint(err)

		s.Equal("some error", err.Error())
	})

	s.Run("returns an ErrAsRequired if the error is a not null constraint of a relation", func() {
		err := errors.New(
			"null value in column \"column\" of relation \"table\" violates not-null constraint",
		)

		err = sql.NotNullConstraint(err)

		s.Error(err)
		s.Equal("table.column cannot be null", err.Error())
	})
}
