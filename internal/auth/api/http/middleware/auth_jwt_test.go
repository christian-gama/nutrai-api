package middleware_test

import (
	gohttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/mock"

	"github.com/christian-gama/nutrai-api/internal/auth/api/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/query"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	qryMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/query"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type AuthSuite struct {
	suite.Suite
}

func TestAuthSuite(t *testing.T) {
	suite.RunUnitTest(t, new(AuthSuite))
}

func (s *AuthSuite) TestHandle() {
	const (
		method = http.MethodGet
		path   = "/test"
	)

	type Mock struct {
		AuthHandler *qryMock.Handler[*query.AuthJwtInput, *query.AuthJwtOutput]
	}

	type Sut struct {
		Sut     middleware.AuthJwt
		Mock    *Mock
		Request *gohttp.Request
		Engine  *gin.Engine
		Context *gin.Context
	}

	makeSut := func() *Sut {
		mock := &Mock{
			AuthHandler: qryMock.NewHandler[*query.AuthJwtInput, *query.AuthJwtOutput](s.T()),
		}

		sut := middleware.NewAuthJwt(mock.AuthHandler)

		// Create a new Gin test context.
		writer := httptest.NewRecorder()
		ctx, engine := gin.CreateTestContext(writer)

		// Inject the middleware into the engine.
		engine.Use(sut.Handle)

		// Create a test request with the method(GET, POST, etc) and the URL path.
		request := httptest.NewRequest(method.String(), path, nil)
		gintest.SetAccessToken(request, value.Token(faker.Jwt()))

		return &Sut{Sut: sut, Mock: mock, Request: request, Engine: engine, Context: ctx}
	}

	s.Run("returns 200 when no error is thrown", func() {
		sut := makeSut()

		sut.Mock.AuthHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(fake.AuthJwtOutput(), nil)

		sut.Engine.Handle(method.String(), path, func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		sut.Engine.ServeHTTP(sut.Context.Writer, sut.Request)

		s.Equal(http.StatusOK, sut.Context.Writer.Status())
	})

	s.Run("returns 401 when there is no token", func() {
		sut := makeSut()

		sut.Request.Header.Del("Authorization")

		sut.Engine.Handle(method.String(), path, func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		sut.Engine.ServeHTTP(sut.Context.Writer, sut.Request)

		s.Equal(http.StatusUnauthorized, sut.Context.Writer.Status())
	})

	s.Run("panics when the token is invalid", func() {
		sut := makeSut()

		sut.Mock.AuthHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(&query.AuthJwtOutput{}, errors.Unauthorized("any error"))

		sut.Engine.Handle(method.String(), path, func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		sut.Engine.ServeHTTP(sut.Context.Writer, sut.Request)

		s.Equal(http.StatusUnauthorized, sut.Context.Writer.Status())
	})
}
