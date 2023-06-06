package controller

import (
	"strings"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/core/infra/validation"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	"github.com/christian-gama/nutrai-api/pkg/unit"
	"github.com/gin-gonic/gin"
)

// Handler is a function that handles a HTTP request.
type Handler[P any] func(*gin.Context, *P)

// Controller is a interface that represents a controller for gin.
// It handles the request and response and binds the data to the input.
// The data can be either a request query string, uri param or a request body.
type Controller interface {
	// Handle is the function that will be called by the router.
	Handle(ctx *gin.Context)

	// Method is the HTTP method that the handler will listen to.
	Method() http.Method

	// Path is the path that the handler will listen to.
	Path() Path

	// Params is the list of params that the handler will listen to.
	Params() Params

	// RPM is the rate limit per minute for the handler. It's used to limit the number of requests
	// per minute for the given endpoint. The default value is 0, which means that the endpoint
	// will not be rate limited, unless the global rate limit is set. The priority is the
	// controller's rate limit, then the global rate limit.
	RPM() int

	// AuthStrategy is the security level of the endpoint. It's used to determine if the endpoint
	// requires authentication or not.
	AuthStrategy() AuthStrategy
}

// ControllerOptions is the options for the controller constructor. It's used
// to setup the controller before using it.
type Options struct {
	// Params is the list of params that the handler will listen to.
	Params Params

	// Path is the path that the handler will listen to.
	Path Path

	// Method is the HTTP method that the handler will listen to.
	Method http.Method

	// RPM is the rate limit per minute for the handler. It's used to limit the number of requests
	// per minute for the given endpoint. The default value is 0, which means that the endpoint
	// will not be rate limited, unless the global rate limit is set. The priority is the
	// controller's rate limit, then the global rate limit.
	RPM int

	// AuthStrategy is the security level of the endpoint. It's used to determine if the endpoint
	// requires authentication or not.
	AuthStrategy AuthStrategy
}

// controllerImpl is the implementation of the Controller interface.
type controllerImpl[Input any] struct {
	handler      func(ctx *gin.Context, input *Input)
	method       http.Method
	path         Path
	rpm          int
	params       Params
	authStrategy AuthStrategy

	input Input
}

// NewController creates a new controller.
func NewController[Input any](
	handler Handler[Input],
	opts Options,
) Controller {
	errutil.MustBeNotEmpty("handler", handler)
	errutil.MustBeNotEmpty("options.Method", opts.Method)

	controller := &controllerImpl[Input]{
		handler:      handler,
		method:       opts.Method,
		path:         opts.Path,
		params:       opts.Params,
		rpm:          opts.RPM,
		authStrategy: opts.AuthStrategy,
	}

	controller.validate()

	return controller
}

// Handle implements Controller.
func (c controllerImpl[Input]) Handle(ctx *gin.Context) {
	http.ExtractBody(ctx, &c.input)
	http.ExtractQuery(ctx, &c.input)
	http.ExtractParams(ctx, &c.input)
	http.ExtractCurrentUser(ctx, &c.input)

	if ctx.IsAborted() {
		return
	}

	response.Response(ctx, func() {
		validator := validation.MakeValidator()

		err := validator.Validate(c.input)
		if err != nil {
			panic(err)
		}

		c.handler(ctx, &c.input)
	})
}

// Handler implements Controller.
func (c *controllerImpl[P]) Method() http.Method {
	return c.method
}

// RPM implements Controller.
func (c *controllerImpl[P]) RPM() int {
	return c.rpm
}

// Path implements Controller.
func (c *controllerImpl[P]) Path() Path {
	return c.path
}

// Params implements Controller.
func (c *controllerImpl[P]) Params() Params {
	return c.params
}

// AuthStrategy implements Controller.
func (c *controllerImpl[P]) AuthStrategy() AuthStrategy {
	return c.authStrategy
}

// validate validates the controller.
func (c *controllerImpl[P]) validate() {
	var result *errutil.Error

	for _, param := range c.params {
		if !c.hasValidCharacters(unit.Alphabet, param) {
			result = errutil.Append(
				result,
				errors.InternalServerError("the param %s contains invalid characters", param),
			)
		}

		if slice.Count(c.params, param) > 1 {
			result = errutil.Append(
				result,
				errors.InternalServerError("the param %s is duplicated", param),
			)
		}
	}

	if !strings.HasPrefix(c.path.String(), "/") {
		result = errutil.Append(
			result,
			errors.InternalServerError("the path %s does not start with a slash", c.path),
		)
	}

	if !c.hasValidCharacters(append(unit.AlphaNumeric, []rune("-/")...), c.path.String()) {
		result = errutil.Append(
			result,
			errors.InternalServerError("the path %s contains invalid characters", c.path),
		)
	}

	methods := []http.Method{
		http.MethodDelete, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch,
	}
	if !slice.Contains(methods, c.method) {
		result = errutil.Append(
			result,
			errors.InternalServerError("the method %s is invalid", c.method),
		)
	}

	if c.rpm < 0 {
		result = errutil.Append(
			result,
			errors.InternalServerError("the rpm %d is invalid", c.rpm),
		)
	}

	if c.authStrategy == nil {
		result = errutil.Append(
			result,
			errors.InternalServerError("the auth strategy is nil"),
		)
	}

	if result.HasErrors() {
		panic(result)
	}
}

func (c *controllerImpl[Input]) hasValidCharacters(allowed []rune, s string) bool {
	for _, r := range s {
		if !slice.Contains(allowed, r) {
			return false
		}
	}

	return true
}
