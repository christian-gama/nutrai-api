package middleware

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Cors = middleware.Middleware

func NewCors() Cors {
	handler := cors.New(cors.Config{
		AllowFiles:    true,
		AllowOrigins:  env.App.AllowedOrigins,
		AllowWildcard: true,
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{
			http.MethodGet.String(),
			http.MethodPost.String(),
			http.MethodPut.String(),
			http.MethodDelete.String(),
			http.MethodPatch.String(),
		},
	})

	return middleware.NewMiddleware(func(ctx *gin.Context) {
		handler(ctx)
		ctx.Next()
	})
}
