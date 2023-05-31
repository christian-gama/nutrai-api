package middleware_test

import (
	gohttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/exception/api/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	cmdMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/command"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type SaveExceptionSuite struct {
	suite.Suite
}

func TestSaveExceptionSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SaveExceptionSuite))
}

func (s *SaveExceptionSuite) TestHandle() {
	const (
		method = http.MethodGet
		path   = "/test"
	)

	type Mock struct {
		SaveExceptionHandler *cmdMock.Handler[*command.SaveExceptionInput]
	}

	type Sut struct {
		Sut     middleware.SaveException
		Mock    *Mock
		Request *gohttp.Request
		Engine  *gin.Engine
		Context *gin.Context
	}

	makeSut := func() *Sut {
		mock := &Mock{
			SaveExceptionHandler: cmdMock.NewHandler[*command.SaveExceptionInput](s.T()),
		}

		sut := middleware.NewSaveException(mock.SaveExceptionHandler)

		// Create a new Gin test context.
		writer := httptest.NewRecorder()
		ctx, engine := gin.CreateTestContext(writer)

		// Inject the middleware into the engine.
		engine.Use(sut.Handle)

		// Create a test request with the method(GET, POST, etc) and the URL path.
		request := httptest.NewRequest(method.String(), path, nil)

		return &Sut{Sut: sut, Mock: mock, Request: request, Engine: engine, Context: ctx}
	}

	s.Run("returns 200 when no error is thrown", func() {
		sut := makeSut()

		sut.Engine.Handle(method.String(), path, func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		sut.Engine.ServeHTTP(sut.Context.Writer, sut.Request)

		s.Equal(http.StatusOK, sut.Context.Writer.Status())
	})

	s.Run("fails and return 500 when an internal error is thrown", func() {
		sut := makeSut()

		sut.Engine.Handle(method.String(), path, func(c *gin.Context) {
			panic(errors.InternalServerError("something went wrong, please try again later"))
		})

		sut.Mock.SaveExceptionHandler.
			On("Handle", sut.Request.Context(), mock.Anything).
			Return(nil)

		sut.Engine.ServeHTTP(sut.Context.Writer, sut.Request)

		s.Equal(http.StatusInternalServerError, sut.Context.Writer.Status())
	})

	s.Run("fails and returns 500 when an errutil.Error different from internal is thrown", func() {
		sut := makeSut()

		sut.Engine.Handle(method.String(), path, func(c *gin.Context) {
			panic(errors.Invalid("test", "test"))
		})

		sut.Mock.SaveExceptionHandler.
			On("Handle", sut.Request.Context(), mock.Anything).
			Return(nil)

		sut.Engine.ServeHTTP(sut.Context.Writer, sut.Request)

		s.Equal(http.StatusInternalServerError, sut.Context.Writer.Status())
	})

	s.Run("fails and returns 500 when something different from error is thrown", func() {
		sut := makeSut()

		sut.Engine.Handle(method.String(), path, func(c *gin.Context) {
			panic("test")
		})

		sut.Mock.SaveExceptionHandler.
			On("Handle", sut.Request.Context(), mock.Anything).
			Return(nil)

		sut.Engine.ServeHTTP(sut.Context.Writer, sut.Request)

		s.Equal(http.StatusInternalServerError, sut.Context.Writer.Status())
	})
}
