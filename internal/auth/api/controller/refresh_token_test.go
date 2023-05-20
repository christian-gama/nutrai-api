package controller_test

import (
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/api/controller"
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	jwtValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/service"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	serviceMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/app/service"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RefreshTokenSuite struct {
	suite.Suite
}

func TestRefreshTokenSuite(t *testing.T) {
	suite.RunUnitTest(t, new(RefreshTokenSuite))
}

func (s *RefreshTokenSuite) TestHandle() {
	type Mock struct {
		RefreshTokenHandler *serviceMock.Handler[*service.RefreshTokenInput, *service.RefreshTokenOutput]
	}

	type Sut struct {
		Sut   controller.RefreshToken
		Input *service.RefreshTokenInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.RefreshTokenInput()
		mocks := &Mock{
			RefreshTokenHandler: serviceMock.NewHandler[*service.RefreshTokenInput, *service.RefreshTokenOutput](
				s.T(),
			),
		}
		sut := controller.NewRefreshToken(mocks.RefreshTokenHandler)
		return &Sut{Sut: sut, Mock: mocks, Input: input}
	}

	s.Run(
		"should return an access token and refresh token when refreshing a token succeeds",
		func() {
			sut := makeSut()

			accessToken := jwtValue.Token("access")
			sut.Mock.RefreshTokenHandler.
				On("Handle", mock.Anything, sut.Input).
				Return(&service.RefreshTokenOutput{
					Access: accessToken,
				}, nil)

			ctx, body := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusOK, ctx.Writer.Status())
			s.EqualValues(
				accessToken,
				body.Data.(map[string]any)["access"],
				"should return access token",
			)
		})

	s.Run("Refresh", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.Refresh = ""

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when invalid", func() {
			sut := makeSut()

			sut.Input.Refresh = "invalid_token"

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("panics when RefreshTokenHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.RefreshTokenHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
