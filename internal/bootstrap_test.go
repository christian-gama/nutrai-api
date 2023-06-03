package internal_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal"
	"github.com/christian-gama/nutrai-api/internal/core/infra/redis/conn"
	sqlconn "github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type BootstrapSuite struct {
	suite.Suite
}

var ran = false

func TestBootstrapSuite(t *testing.T) {
	suite.RunUnitTestOnce(t, new(BootstrapSuite), &ran)
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
}
