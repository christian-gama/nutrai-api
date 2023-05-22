package http

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/auth/infra/store"
	"github.com/gin-gonic/gin"
)

func extract[Input any](ctx *gin.Context, input *Input, extractFn func(any) error) error {
	if ctx.IsAborted() {
		return nil
	}

	if reflect.TypeOf(input).Kind() == reflect.Ptr &&
		reflect.TypeOf(input).Elem().Kind() == reflect.Interface {
		return nil
	}

	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Ptr {
		panic(errors.New("input must be a pointer"))
	}

	if v.IsNil() {
		panic(errors.New("input is nil"))
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
		BadRequest(ctx, fmt.Errorf("could not extract body: %w", err))
	}
}

// ExtractQuery extracts the query from the query string.
func ExtractQuery[Query any](ctx *gin.Context, query *Query) {
	err := extract(ctx, query, ctx.ShouldBindQuery)
	if err != nil {
		BadRequest(ctx, fmt.Errorf("could not extract query: %w", err))
	}
}

// ExtractParams extracts the params from the request URI.
func ExtractParams[Params any](ctx *gin.Context, params *Params) {
	err := extract(ctx, params, ctx.ShouldBindUri)
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		BadRequest(ctx, fmt.Errorf("could not extract params: %w", err))
	}
}

// ExtractCurrentUser extracts the current user from the request. It expects the input to have a
// field with the tag `ctx:"currentUser"`.
func ExtractCurrentUser[Data any](ctx *gin.Context, data *Data) {
	if ctx.IsAborted() {
		return
	}

	v := reflect.ValueOf(data).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("ctx")

		if tag == "currentUser" {
			currentUser, err := store.GetUser(ctx)
			if err != nil {
				field.Set(reflect.Zero(field.Type()))
				InternalServerError(ctx, err)
			}

			field.Set(reflect.ValueOf(currentUser))
		}
	}
}
