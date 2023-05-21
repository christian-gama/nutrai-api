package persistence_test

import (
	"context"
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/exception/infra/persistence/sql"
	"github.com/christian-gama/nutrai-api/internal/exception/infra/persistence/sql/schema"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/exception/domain/model/exception"
	fixture "github.com/christian-gama/nutrai-api/testutils/fixture/exception/sql"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"gorm.io/gorm"
)

type ExceptionSuite struct {
	suite.SuiteWithSQLConn
	Exception func(db *gorm.DB) repo.Exception
}

func TestExceptionSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(ExceptionSuite))
}

func (s *ExceptionSuite) SetupTest() {
	s.Exception = func(db *gorm.DB) repo.Exception {
		return persistence.NewSQLException(db)
	}
}

func (s *ExceptionSuite) TestSave() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.SaveExceptionInput) (*exception.Exception, error)
		Ctx   context.Context
		Input repo.SaveExceptionInput
	}

	makeSut := func(db *gorm.DB) Sut {
		exception := fake.Exception()

		input := repo.SaveExceptionInput{
			Exception: exception,
		}

		sut := s.Exception(db).Save

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
		}
	}

	s.Run("Should create a new exception", func(db *gorm.DB) {
		sut := makeSut(db)

		exception, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.NotZero(exception.ID, "Should have an ID")
	})
}

func (s *ExceptionSuite) TestDeleteOld() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.DeleteOldExceptionInput) error
		Ctx   context.Context
		Input repo.DeleteOldExceptionInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.DeleteOldExceptionInput{}
		sut := s.Exception(db).DeleteOld

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should delete an exception that is old", func(db *gorm.DB) {
		sut := makeSut(db)

		fixture.SaveException(db, nil)

		sut.Input.BeforeDate = time.Now().Add(-time.Hour * 24 * 6)

		err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.SQLRecordDoesNotExist(db, &schema.Exception{})
	})

	s.Run("Should not delete an exception that is fresh", func(db *gorm.DB) {
		sut := makeSut(db)

		fixture.SaveException(db, nil)

		sut.Input.BeforeDate = time.Now().Add(time.Hour * 24 * 7)

		err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.SQLCount(db, &schema.Exception{}, 1)
	})

	s.Run("Should delete nothing if the exception does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.BeforeDate = time.Now().Add(-time.Hour * 24 * 6)

		err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.SQLRecordDoesNotExist(db, &schema.Exception{})
	})
}
