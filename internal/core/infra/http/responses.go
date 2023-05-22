package http

import (
	"net/http"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/gin-gonic/gin"
)

// InternalServerError is a helper function to return an internal server error.
// It will panic and be recovered by the middleware from errors module.
func InternalServerError(ctx *gin.Context, err error) {
	if env.Config.Debug {
		log.WithCaller.Warnf("Internal Server Error: %s", err.Error())
	}

	ctx.Errors = append(ctx.Errors, ctx.Error(err))
	panic(err)
}

// BadRequest is a helper function to return a bad request error
// with a default JSON response.
func BadRequest(ctx *gin.Context, err error) {
	if env.Config.Debug {
		log.WithCaller.Warnf("Bad Request: %s", err.Error())
	}

	ctx.Errors = append(ctx.Errors, ctx.Error(err))
	ctx.AbortWithStatusJSON(http.StatusBadRequest, Error(err))
}

// NotFound is a helper function to return a not found error
// with a default JSON response.
func NotFound(ctx *gin.Context, err error) {
	if env.Config.Debug {
		log.WithCaller.Warnf("Not Found: %s", err.Error())
	}

	ctx.Errors = append(ctx.Errors, ctx.Error(err))
	ctx.AbortWithStatusJSON(http.StatusNotFound, Error(err))
}

// Created is a helper function to return a created response
// with a default JSON response.
func Created(ctx *gin.Context, data any) {
	ctx.AbortWithStatusJSON(http.StatusCreated, Data(data))
}

// Ok is a helper function to return a ok response with a
// default JSON response.
func Ok(ctx *gin.Context, data any) {
	ctx.AbortWithStatusJSON(http.StatusOK, Data(data))
}

// NoContent is a helper function to return a no content response.
func NoContent(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNoContent)
}
