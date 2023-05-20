package http

import (
	"errors"
	"fmt"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/core/infra/validation"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	"github.com/christian-gama/nutrai-api/pkg/unit"
	"github.com/gin-gonic/gin"
)

type Handler[P any] func(*gin.Context, *P)

// Controller is a interface that represents a controller for gin.
// It handles the request and response and binds the data to the input.
// The data can be either a request query string, uri param or a request body.
type Controller interface {
	// Handle is the function that will be called by the router.
	Handle(ctx *gin.Context)

	// Method is the HTTP method that the handler will listen to.
	Method() Method

	// Path is the path that the handler will listen to.
	Path() Path

	// Params is the list of params that the handler will listen to.
	Params() Params

	// IsPublic is a flag that indicates if the handler is public or not.
	IsPublic() bool
}

// ControllerOptions is the options for the controller constructor. It's used
// to setup the controller before using it.
type ControllerOptions struct {
	IsPublic bool
	Params   Params
	Path     Path
	Method   Method
}

// controllerImpl is the implementation of the Controller interface.
type controllerImpl[Input any] struct {
	handler  func(*gin.Context, *Input)
	method   Method
	path     Path
	isPublic bool
	params   Params

	input Input
}

// NewController creates a new controller.
func NewController[Input any](
	handler Handler[Input],
	opts ControllerOptions,
) Controller {
	if handler == nil {
		panic(errors.New("handler is nil"))
	}

	if opts.Method == "" {
		panic(errors.New("method is empty"))
	}

	controller := &controllerImpl[Input]{
		handler:  handler,
		method:   opts.Method,
		path:     opts.Path,
		isPublic: opts.IsPublic,
		params:   opts.Params,
	}

	controller.validate()

	return controller
}

// Handle implements Controller.
func (c controllerImpl[Input]) Handle(ctx *gin.Context) {
	ExtractBody(ctx, &c.input)
	ExtractQuery(ctx, &c.input)
	ExtractParams(ctx, &c.input)

	if ctx.IsAborted() {
		return
	}

	Response(ctx, func() {
		validator := validation.MakeValidator()

		err := validator.Validate(c.input)
		if err != nil {
			panic(err)
		}

		c.handler(ctx, &c.input)
	})
}

// Handler implements Controller.
func (c *controllerImpl[P]) Method() Method {
	return c.method
}

// Path implements Controller.
func (c *controllerImpl[P]) Path() Path {
	return c.path
}

// IsPublic implements Controller.
func (c *controllerImpl[P]) IsPublic() bool {
	return c.isPublic
}

// Params implements Controller.
func (c *controllerImpl[P]) Params() Params {
	return c.params
}

// validate validates the controller.
func (c *controllerImpl[P]) validate() {
	var result *errutil.Error

	for _, param := range c.params {
		if !hasValidCharacters(unit.Alphabet, param) {
			result = errutil.Append(
				result,
				fmt.Errorf("the param %s contains invalid characters", param),
			)
		}

		if slice.Count(c.params, param) > 1 {
			result = errutil.Append(
				result,
				fmt.Errorf("the param %s is duplicated", param),
			)
		}
	}

	if !strings.HasPrefix(c.path.String(), "/") {
		result = errutil.Append(
			result,
			fmt.Errorf("the path %s does not start with a slash", c.path),
		)
	}

	if !hasValidCharacters(append(unit.AlphaNumeric, []rune("-/")...), c.path.String()) {
		result = errutil.Append(
			result,
			fmt.Errorf("the path %s contains invalid characters", c.path),
		)
	}

	methods := []Method{
		MethodDelete, MethodGet, MethodPost, MethodPut,
	}

	if !slice.Contains(methods, c.method) {
		panic(fmt.Errorf("the method %s is invalid", c.method))
	}

	if result.HasErrors() {
		panic(result)
	}
}

func hasValidCharacters(allowed []rune, s string) bool {
	for _, r := range s {
		if !slice.Contains(allowed, r) {
			return false
		}
	}

	return true
}
