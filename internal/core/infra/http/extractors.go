package http

import (
	"reflect"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/auth/infra/ctxstore"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/christian-gama/nutrai-api/pkg/reflection"
	"github.com/gin-gonic/gin"
)

// extract is a generic function to extract data from a context into an input structure. The actual
// extraction is performed by the provided extractFn function.
func extract[Input any](ctx *gin.Context, input *Input, extractFn func(any) error) error {
	if err := validateInput(ctx, input); err != nil {
		panic(err)
	}

	err := extractFn(input)
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		return err
	}

	return nil
}

// ExtractBody extracts the body from the request body.
func ExtractBody[Body any](ctx *gin.Context, body *Body) {
	err := extract(ctx, body, ctx.ShouldBindJSON)
	if err != nil {
		response.BadRequest(ctx, errors.Invalid("could not extract body: %s", err.Error()))
	}
}

// ExtractQuery extracts the query from the query string.
func ExtractQuery[Query any](ctx *gin.Context, query *Query) {
	err := extract(ctx, query, ctx.ShouldBindQuery)
	if err != nil {
		response.BadRequest(ctx, errors.Invalid("could not extract query: %s", err.Error()))
	}
}

// ExtractParams extracts the params from the request URI.
func ExtractParams[Params any](ctx *gin.Context, params *Params) {
	err := extract(ctx, params, ctx.ShouldBindUri)
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		response.BadRequest(ctx, errors.Invalid("could not extract params: %s", err.Error()))
	}
}

// ExtractCurrentUser extracts the current user from the request. It expects the input to have a
// field with the tag `ctx:"current_user"`.
func ExtractCurrentUser[Input any](ctx *gin.Context, input *Input) {
	if err := validateInput(ctx, input); err != nil {
		panic(err)
	}

	if reflection.IsPointerToInterface(input) {
		return
	}

	reflection.IterateStructFields(input, func(opts *reflection.FieldIterationOptions) {
		if opts.Tag.Get("ctx") == "current_user" {
			currentUser, err := ctxstore.GetUser(ctx)
			if err != nil {
				opts.Field.Set(reflect.Zero(opts.Field.Type()))
				response.InternalServerError(ctx, err)
			}

			opts.Field.Set(reflect.ValueOf(currentUser))
		}
	})
}

func validateInput[T any](ctx *gin.Context, input T) error {
	if ctx.IsAborted() {
		return nil
	}

	if reflection.IsPointerToInterface(input) {
		return nil
	}

	if !reflection.IsPointer(input) {
		return errors.InternalServerError("input must be a pointer")
	}

	if reflection.IsNil(input) {
		return errors.InternalServerError("input cannot be nil")
	}

	return nil
}
