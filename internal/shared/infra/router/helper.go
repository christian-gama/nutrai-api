package router

import (
	"github.com/christian-gama/nutrai-api/pkg/log"
	"github.com/fatih/color"
)

// logLevel returns a log function based on the status code.
func logLevel(status int) func(string, ...interface{}) {
	switch {
	case status >= 500:
		return log.New(&log.Config{}).Errorf

	default:
		return log.New(&log.Config{}).Infof
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
