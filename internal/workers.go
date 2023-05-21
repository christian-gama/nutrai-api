package internal

import (
	"github.com/christian-gama/nutrai-api/internal/exception"
)

// Workers registers all workers.
func Workers() {
	exception.Workers()
}
