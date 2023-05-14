package internal_test

import (
	"context"
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/internal"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/shared/domain/logger"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/mock"
)

type BootstrapSuite struct {
	suite.Suite
}

func TestBootstrapSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(BootstrapSuite))
}

func (s *BootstrapSuite) TestBootstrap() {
	s.Run("do not panic", func() {
		log := mocks.NewLogger(s.T())
		log.On("Infof", mock.Anything, mock.Anything)
		log.On("Infof", mock.Anything, mock.Anything, mock.Anything)
		log.On("Infof", mock.Anything, mock.Anything)
		log.On("Fatal", mock.Anything).Maybe()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
		defer cancel()

		s.NotPanics(func() { internal.Bootstrap(ctx, log, ".env.test") })
	})
}
