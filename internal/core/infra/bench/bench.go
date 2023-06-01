package bench

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

// Duration returns the duration that a function takes to execute.
func Duration(fn func()) time.Duration {
	start := time.Now()
	fn()
	elapsed := time.Since(start)

	return elapsed
}

// PrintDuration prints the duration that a function takes to execute.
func PrintDuration(resource string, fn func()) {
	duration := Duration(fn)
	log.Infof("%s took %s to complete", resource, duration)
}
