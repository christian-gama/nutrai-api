package http

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func extract[Payload any](ctx *gin.Context, payload *Payload, extractFn func(any) error) error {
	if ctx.IsAborted() {
		return nil
	}

	if reflect.TypeOf(payload).Kind() == reflect.Ptr &&
		reflect.TypeOf(payload).Elem().Kind() == reflect.Interface {
		return nil
	}

	v := reflect.ValueOf(payload)
	if v.Kind() != reflect.Ptr {
		panic(errors.New("payload must be a pointer"))
	}

	if v.IsNil() {
		panic(errors.New("payload is nil"))
	}

	err := extractFn(payload)
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		return err
	}

	return nil
}

// ExtractBody extracts the body from the request body.
func ExtractBody[Body any](ctx *gin.Context, payload *Body) {
	err := extract(ctx, payload, ctx.ShouldBindJSON)
	if err != nil {
		BadRequest(ctx, fmt.Errorf("could not extract body: %w", err))
	}
}

// ExtractQuery extracts the query from the query string.
func ExtractQuery[Query any](ctx *gin.Context, payload *Query) {
	err := extract(ctx, payload, ctx.ShouldBindQuery)
	if err != nil {
		BadRequest(ctx, fmt.Errorf("could not extract query: %w", err))
	}
}

// ExtractParams extracts the params from the request URI.
func ExtractParams[Params any](ctx *gin.Context, payload *Params) {
	err := extract(ctx, payload, ctx.ShouldBindUri)
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		BadRequest(ctx, fmt.Errorf("could not extract params: %w", err))
	}
}
