package internal_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/christian-gama/nutrai-api/internal"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/redis/conn"
	sqlconn "github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	"github.com/stretchr/testify/suite"
)

type BootstrapSuite struct {
	suite.Suite
}

var ran = false

func TestBootstrapSuite(t *testing.T) {
	if !ran {
		ran = true
		modes := []string{"unit", "integration", "all"}
		mode, ok := os.LookupEnv("TEST_MODE")
		if !ok {
			t.Fatal("TEST_MODE is not set")
		}

		if !slice.Contains(modes, mode) {
			panic(fmt.Errorf("expected TEST_MODE to be one of: %v", modes))
		}

		if mode != "unit" {
			suite.Run(t, new(BootstrapSuite))
		} else {
			t.SkipNow()
		}
	}
}

func (s *BootstrapSuite) TestBootstrap() {
	s.Run("Should start without panicking", func() {
		s.NotPanics(func() {
			internal.Bootstrap(".env.test")
		})
	})

	s.Run("Should have postgres connection", func() {
		psql := sqlconn.GetPsql()
		s.NotNil(psql)
	})

	s.Run("Should have redis connection", func() {
		redis := conn.GetRedis()
		s.NotNil(redis)
	})

	s.Run("Should have jwt security middleware set", func() {
		s.NotNil(
			controller.SecurityJwt.Middleware(),
			"You must provide a middleware for security jwt",
		)
	})

	s.Run("Should have api key security middleware set", func() {
		s.NotNil(
			controller.SecurityApiKey.Middleware(),
			"You must provide a middleware for security api_key",
		)
	})
}
