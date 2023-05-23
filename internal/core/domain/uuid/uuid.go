package uuid

import "github.com/christian-gama/nutrai-api/internal/core/domain/value"

// Generator is the interface that wraps the Generate method.
type Generator interface {
	// Generate generates a new UUID.
	Generate() value.UUID
}
