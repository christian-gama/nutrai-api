package sqlerr

import (
	"errors"
	"testing"

	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type ConstraintsSuite struct {
	suite.Suite
}

func TestConstraintsSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ConstraintsSuite))
}

func (s *ConstraintsSuite) TestCheckConstraint() {
	s.Run("returns an ErrInvalid if the error is a check constraint violation", func() {
		err := errors.New("chk__column__message")

		err = checkConstraint(err)

		s.ErrorAsInvalid(err)
	})

	s.Run("returns the original error if it is not a check constraint violation", func() {
		err := errors.New("some error")

		err = checkConstraint(err)

		s.Equal("some error", err.Error())
	})
}

func (s *ConstraintsSuite) TestForeignKeyConstraint() {
	s.Run("returns an ErrNotFound if the error is a foreign key constraint violation", func() {
		err := errors.New(
			`"table" violates foreign key constraint "fk__column__refTable.refColumn"`,
		)

		err = foreignKeyConstraint(err)

		s.ErrorAsNotFound(err)
	})

	s.Run("returns the original error if it is not a foreign key constraint violation", func() {
		err := errors.New("some error")

		err = foreignKeyConstraint(err)

		s.Equal("some error", err.Error())
	})
}

func (s *ConstraintsSuite) TestNotNullConstraint() {
	s.Run("returns an ErrRequired if the error is a not null constraint violation", func() {
		err := errors.New("null value in column \"column\" violates not-null constraint")

		err = notNullConstraint(err)

		s.ErrorAsRequired(err)
	})

	s.Run("returns the original error if it is not a not null constraint violation", func() {
		err := errors.New("some error")

		err = notNullConstraint(err)

		s.Equal("some error", err.Error())
	})

	s.Run("returns an ErrAsRequired if the error is a not null constraint of a relation", func() {
		err := errors.New(
			"null value in column \"column\" of relation \"table\" violates not-null constraint",
		)

		err = notNullConstraint(err)

		s.ErrorAsRequired(err)
	})
}

func (s *ConstraintsSuite) TestUniqueConstraint() {
	s.Run("returns an ErrAlreadyExists if the error is a unique constraint violation", func() {
		err := errors.New("uidx__table__column")

		err = uniqueConstraint(err)

		s.ErrorAsAlreadyExists(err)
	})

	s.Run("returns the original error if it is not a unique constraint violation", func() {
		err := errors.New("some error")

		err = uniqueConstraint(err)

		s.Equal("some error", err.Error())
	})

	s.Run(
		"returns an ErrAlreadyExists if the error is a unique constraint violation but comes from chk",
		func() {
			err := errors.New("chk__column__message")

			err = uniqueConstraint(err)

			s.ErrorAsAlreadyExists(err)
		},
	)
}
