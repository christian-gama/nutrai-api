package suite

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/env"
	"github.com/christian-gama/nutrai-api/testutils/faker"
)

func init() {
	faker.InitializeProviders()
	faker.Setup()

	env.Load(".env.test")
}
