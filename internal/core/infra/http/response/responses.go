package response

import (
	"net/http"

	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/gin-gonic/gin"
)

// InternalServerError is a helper function to return an internal server error. It will panic and be
// recovered by the middleware from errors module.
func InternalServerError(ctx *gin.Context, err error) {
	log := log.MakeWithCaller()
	log.Debugf("Internal Server Error: %s", err.Error())

	ctx.Errors = append(ctx.Errors, ctx.Error(err))

	panic(err)
}

// BadRequest is a helper function to return a bad request error with a default JSON response.
func BadRequest(ctx *gin.Context, err error) {
	log := log.MakeWithCaller()
	log.Debugf("Bad Request: %s", err.Error())

	ctx.Errors = append(ctx.Errors, ctx.Error(err))
	ctx.AbortWithStatusJSON(http.StatusBadRequest, Error(err))
}

// Unauthorized is a helper function to return an unauthorized error with a default JSON response.
func Unauthorized(ctx *gin.Context, err error) {
	log := log.MakeWithCaller()
	log.Debugf("Unauthorized: %s", err.Error())

	ctx.Errors = append(ctx.Errors, ctx.Error(err))
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, Error(err))
}

// NotFound is a helper function to return a not found error with a default JSON response.
func NotFound(ctx *gin.Context, err error) {
	log := log.MakeWithCaller()
	log.Debugf("Not Found: %s", err.Error())

	ctx.Errors = append(ctx.Errors, ctx.Error(err))
	ctx.AbortWithStatusJSON(http.StatusNotFound, Error(err))
}

// Created is a helper function to return a created response
// with a default JSON response.
func Created(ctx *gin.Context, data any) {
	log := log.MakeWithCaller()
	log.Debugf("Created: %v", data)

	ctx.AbortWithStatusJSON(http.StatusCreated, Data(data))
}

// Ok is a helper function to return a ok response with a default JSON response.
func Ok(ctx *gin.Context, data any) {
	log := log.MakeWithCaller()
	log.Debugf("Ok: %v", data)

	ctx.AbortWithStatusJSON(http.StatusOK, Data(data))
}

// NoContent is a helper function to return a no content response.
func NoContent(ctx *gin.Context) {
	log := log.MakeWithCaller()
	log.Debugf("No Content")

	ctx.AbortWithStatus(http.StatusNoContent)
}

// TooManyRequests is a helper function to return a too many requests error with a default JSON
// response.
func TooManyRequests(ctx *gin.Context, err error) {
	log := log.MakeWithCaller()
	log.Debugf("Too Many Requests: %s", err.Error())

	ctx.Errors = append(ctx.Errors, ctx.Error(err))
	ctx.AbortWithStatusJSON(http.StatusTooManyRequests, Error(err))
}
