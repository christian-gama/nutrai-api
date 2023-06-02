package middleware

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/infra/bench"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

type Logging = middleware.Middleware

func NewLogging() Logging {
	return middleware.NewMiddleware(NewLoggingHandler().Handle)
}

type loggingHandlerImpl struct{}

func NewLoggingHandler() *loggingHandlerImpl {
	return &loggingHandlerImpl{}
}

func (l *loggingHandlerImpl) Handle(ctx *gin.Context) {
	duration := bench.Duration(ctx.Next)

	l.logLevel(ctx.Writer.Status(), duration)(
		"%-6s | %-5s | %6s | %s",
		ctx.Request.Method,
		l.statusColor(ctx.Writer.Status()),
		duration.Truncate(time.Millisecond),
		ctx.Request.URL.Path,
	)
}

// logLevel returns a log function based on the status code.
func (*loggingHandlerImpl) logLevel(
	status int,
	duration time.Duration,
) func(template string, args ...any) {
	switch {
	case status >= 500:
		return log.Errorf

	default:
		if duration > 500*time.Millisecond {
			return log.Warnf
		}

		return log.Infof
	}
}

// statusColor returns a color based on the status code.
func (*loggingHandlerImpl) statusColor(status int) string {
	var cl *color.Color
	switch {
	case status >= 500:
		cl = color.New(color.FgHiWhite, color.BgRed)

	case status >= 400:
		cl = color.New(color.FgHiWhite, color.BgYellow)

	default:
		cl = color.New(color.FgHiWhite, color.BgGreen)
	}

	return cl.Sprintf(" %d ", status)
}
