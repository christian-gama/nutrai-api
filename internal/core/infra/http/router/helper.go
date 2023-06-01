package router

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/fatih/color"
)

// logLevel returns a log function based on the status code.
func logLevel(status int, duration time.Duration) func(string, ...any) {
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
func statusColor(status int) string {
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
