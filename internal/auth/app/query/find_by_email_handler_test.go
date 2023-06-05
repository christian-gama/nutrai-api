package query_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/query"
	userFake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	userRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/repo"
)

type FindByEmailHandlerSuite struct {
	suite.Suite
}

func TestFindByEmailHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindByEmailHandlerSuite))
}

func (s *FindByEmailHandlerSuite) TestUserHandler() {
	type Mock struct {
		UserRepo *userRepoMock.User
	}

	type Sut struct {
		Sut   query.FindByEmailHandler
		Ctx   context.Context
		Input *query.FindByEmailInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			UserRepo: userRepoMock.NewUser(s.T()),
		}

		input := fake.FindByEmailInput()

		sut := query.NewFindByEmailHandler(mock.UserRepo)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should return a UserOutput", func() {
		sut := makeSut()

		user := userFake.User()
		user.Email = sut.Input.Email
		sut.Mock.UserRepo.On("FindByEmail", sut.Ctx, mock.Anything).Return(user, nil)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Require().NotNil(output)
		s.Equal(sut.Input.Email, output.Email, "Email should be the same")
	})

	s.Run("Should return an error when the repository fails", func() {
		sut := makeSut()

		sut.Mock.UserRepo.On("FindByEmail", sut.Ctx, mock.Anything).Return(nil, assert.AnError)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Error(err)
		s.Nil(output)
		s.ErrorIs(assert.AnError, err)
	})
}
