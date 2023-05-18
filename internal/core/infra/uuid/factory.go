package uuid

import "github.com/christian-gama/nutrai-api/internal/core/domain/uuid"

func MakeGenerator() uuid.Generator {
	return NewGenerator()
}
