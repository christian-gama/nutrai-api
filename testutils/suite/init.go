package suite

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/testutils/faker"
)

func init() {
	faker.InitializeProviders()
	faker.Setup()

	env.NewLoader(".env.test").Load()
}
