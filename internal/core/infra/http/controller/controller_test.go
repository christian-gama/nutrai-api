package controller_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/gin-gonic/gin"
)

type ControllerSuite struct {
	suite.Suite
}

func TestControllerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ControllerSuite))
}

func (s *ControllerSuite) TestNewController() {
	makeSut := func() (
		func(handler controller.Handler[any], opts controller.Options) controller.Controller,
		func(*gin.Context, *any),
		controller.Options,
	) {
		handler := func(*gin.Context, *any) {}
		opts := controller.Options{
			Method:   http.MethodGet,
			Path:     controller.JoinPath(""),
			IsPublic: true,
			Params:   controller.AddParams("id"),
		}
		sut := controller.NewController[any]

		return sut, handler, opts
	}

	s.Run("do not panic when all args are valid", func() {
		sut, handler, opts := makeSut()

		s.NotPanics(func() { sut(handler, opts) })
	})

	s.Run("panic when handler is nil", func() {
		sut, _, opts := makeSut()

		s.Panics(func() { sut(nil, opts) })
	})

	s.Run("panic when method is empty", func() {
		sut, handler, opts := makeSut()

		opts.Method = ""

		s.Panics(func() { sut(handler, opts) })
	})

	s.Run("panic when path is empty", func() {
		sut, handler, opts := makeSut()

		opts.Path = ""

		s.Panics(func() { sut(handler, opts) })
	})

	s.Run("panic when params are invalid", func() {
		sut, handler, opts := makeSut()

		opts.Params = controller.Params{"invalid param"}

		s.Panics(func() { sut(handler, opts) })
	})

	s.Run("panic when params are not unique", func() {
		sut, handler, opts := makeSut()

		opts.Params = controller.AddParams("id").Add("id")

		s.Panics(func() { sut(handler, opts) })
	})

	s.Run("panic when path is invalid", func() {
		sut, handler, opts := makeSut()

		opts.Path = "invalid path"
		s.Panics(func() { sut(handler, opts) })

		opts.Path = "/invalid/:path"
		s.Panics(func() { sut(handler, opts) })
	})

	s.Run("panic when method is invalid", func() {
		sut, handler, opts := makeSut()

		opts.Method = "invalid method"

		s.Panics(func() { sut(handler, opts) })
	})
}

func (s *ControllerSuite) TestController() {
	sut := controller.NewController(
		func(*gin.Context, *any) {},

		controller.Options{
			Method:   http.MethodPut,
			Path:     controller.JoinPath(""),
			IsPublic: true,
			Params:   controller.AddParams("id"),
		},
	)

	s.Run("controller.Method returns the correct method", func() {
		sut := sut.Method

		method := sut()

		s.EqualValues(http.MethodPut, method)
	})

	s.Run("controller.Path returns the correct path", func() {
		sut := sut.Path

		path := sut()

		s.EqualValues("/", path)
	})

	s.Run("controller.IsPublic returns the correct value", func() {
		sut := sut.IsPublic

		isPublic := sut()

		s.True(isPublic)
	})

	s.Run("controller.Params returns the correct params", func() {
		sut := sut.Params

		params := sut()

		s.Equal([]string{"id"}, params.Slice())
	})
}

type Input struct {
	Name string `json:"name" validate:"max=10"`
	ID   int    `uri:"id"    validate:"gte=1"`
	Age  int    `form:"age"  validate:"lte=100"`
}

func (s *ControllerSuite) TestHandle() {
	makeSut := func() (controller.Controller, *Input) {
		input := &Input{
			Name: "John Doe",
			ID:   1,
			Age:  20,
		}

		handler := func(c *gin.Context, p *Input) {
			s.Equal(input.Name, p.Name)
			s.Equal(input.ID, p.ID)
			s.Equal(input.Age, p.Age)
		}

		controller := controller.NewController(
			handler,
			controller.Options{
				Method:   http.MethodPut,
				Path:     controller.JoinPath(""),
				IsPublic: true,
				Params:   controller.AddParams("id"),
			},
		)

		return controller, input
	}

	s.Run("returns 200 when all params are valid", func() {
		sut, input := makeSut()

		ctx := gintest.MustRequest(sut, gintest.Option{
			Data:    map[string]any{"name": input.Name},
			Params:  []string{fmt.Sprintf("%d", input.ID)},
			Queries: fmt.Sprintf("age=%d", input.Age),
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
	})

	s.Run("returns 400 when params are invalid", func() {
		sut, input := makeSut()
		input.Name = "this is an invalid name"

		ctx := gintest.MustRequest(sut, gintest.Option{
			Data:    map[string]any{"name": input.Name},
			Params:  []string{fmt.Sprintf("%d", input.ID)},
			Queries: fmt.Sprintf("age=%d", input.Age),
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("returns 400 when params are not valid", func() {
		sut, input := makeSut()
		input.ID = -1

		ctx := gintest.MustRequest(sut, gintest.Option{
			Data:    map[string]any{"name": input.Name},
			Params:  []string{fmt.Sprintf("%d", input.ID)},
			Queries: fmt.Sprintf("age=%d", input.Age),
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("returns 400 when query is not valid", func() {
		sut, input := makeSut()
		input.Age = 101

		ctx := gintest.MustRequest(sut, gintest.Option{
			Data:    map[string]any{"name": input.Name},
			Params:  []string{fmt.Sprintf("%d", input.ID)},
			Queries: fmt.Sprintf("age=%d", input.Age),
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("returns 200 even when query is not provided", func() {
		sut, input := makeSut()
		input.Age = 0

		ctx := gintest.MustRequest(sut, gintest.Option{
			Data:   map[string]any{"name": input.Name},
			Params: []string{fmt.Sprintf("%d", input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
	})
}
