package bench

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
)

// Duration returns the duration that a function takes to execute.
func Duration(fn func()) time.Duration {
	start := time.Now()
	fn()
	elapsed := time.Since(start)

	return elapsed
}

// PrintDuration prints the duration that a function takes to execute.
func PrintDuration(l logger.Logger, resource string, fn func()) {
	duration := Duration(fn)
	l.Infof("%s took %dms to complete", resource, duration.Milliseconds())
}
