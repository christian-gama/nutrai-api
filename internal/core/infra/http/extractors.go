package http

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/auth/infra/store"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/reflection"
	"github.com/gin-gonic/gin"
)

// extract is a generic function to extract data from a context into an input structure. The actual
// extraction is performed by the provided extractFn function.
func extract[Input any](ctx *gin.Context, input *Input, extractFn func(any) error) error {
	if ctx.IsAborted() {
		return nil
	}

	if reflection.IsPointerToInterface(input) {
		return nil
	}

	if !reflection.IsPointer(input) {
		panic(errors.New("input must be a pointer"))
	}

	if reflection.IsNil(input) {
		panic(errors.New("input cannot be nil"))
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
		response.BadRequest(ctx, fmt.Errorf("could not extract body: %w", err))
	}
}

// ExtractQuery extracts the query from the query string.
func ExtractQuery[Query any](ctx *gin.Context, query *Query) {
	err := extract(ctx, query, ctx.ShouldBindQuery)
	if err != nil {
		response.BadRequest(ctx, fmt.Errorf("could not extract query: %w", err))
	}
}

// ExtractParams extracts the params from the request URI.
func ExtractParams[Params any](ctx *gin.Context, params *Params) {
	err := extract(ctx, params, ctx.ShouldBindUri)
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		response.BadRequest(ctx, fmt.Errorf("could not extract params: %w", err))
	}
}

// ExtractCurrentUser extracts the current user from the request. It expects the input to have a
// field with the tag `ctx:"currentUser"`.
func ExtractCurrentUser[Data any](ctx *gin.Context, data *Data) {
	if ctx.IsAborted() {
		return
	}

	reflection.IterateStructFields(data, func(opts *reflection.FieldIterationOptions) {
		if opts.Tag.Get("ctx") == "currentUser" {
			currentUser, err := store.GetUser(ctx)
			if err != nil {
				opts.Field.Set(reflect.Zero(opts.Field.Type()))
				response.InternalServerError(ctx, err)
			}

			opts.Field.Set(reflect.ValueOf(currentUser))
		}
	})
}
