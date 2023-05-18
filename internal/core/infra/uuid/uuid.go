package uuid

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/uuid"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	_uuid "github.com/google/uuid"
)

// generatorImpl is the implementation of the Generator interface.
type generatorImpl struct{}

// Generate implements uuid.Generator.
func (g *generatorImpl) Generate() value.UUID {
	return value.UUID(_uuid.New().String())
}

// NewGenerator returns a new instance of the Generator interface.
func NewGenerator() uuid.Generator {
	return &generatorImpl{}
}
