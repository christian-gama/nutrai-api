package service_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/value"
	"github.com/christian-gama/nutrai-api/internal/user/app/service"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/service"
	hasherMock "github.com/christian-gama/nutrai-api/testutils/mocks/user/domain/hasher"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
)

type HashPasswordHandlerSuite struct {
	suite.Suite
}

func TestHashPasswordHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(HashPasswordHandlerSuite))
}

func (s *HashPasswordHandlerSuite) TestHandle() {
	type Mock struct {
		Hasher *hasherMock.Hasher
	}

	type Sut struct {
		Sut   func(ctx context.Context, input *service.HashPasswordInput) (*service.HashPasswordOutput, error)
		Input *service.HashPasswordInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		mocks := &Mock{
			Hasher: hasherMock.NewHasher(s.T()),
		}

		sut := service.NewHashPasswordHandler(
			mocks.Hasher,
		)

		input := fake.HashPasswordInput()

		return &Sut{
			Sut:   sut.Handle,
			Mock:  mocks,
			Input: input,
		}
	}

	s.Run("should return a hashed password on success", func() {
		sut := makeSut()

		hashed := value.Password("hashed")
		sut.Mock.Hasher.
			On("Hash", sut.Input.Password).
			Return(hashed, nil)

		output, err := sut.Sut(context.Background(), sut.Input)

		s.Nil(err)
		s.EqualValues(hashed, output.Password, "password should be the same as the hashed")
	})

	s.Run("should return an error if hasher returns an error", func() {
		sut := makeSut()

		sut.Mock.Hasher.
			On("Hash", sut.Input.Password).
			Return(value.Password(""), assert.AnError)

		output, err := sut.Sut(context.Background(), sut.Input)

		s.Nil(output)
		s.ErrorIs(err, assert.AnError)
	})
}
