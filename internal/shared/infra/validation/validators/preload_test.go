package validators_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/shared/infra/validation/validators"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type PreloadSuite struct {
	suite.Suite
}

func TestPreloadSuite(t *testing.T) {
	suite.RunUnitTest(t, new(PreloadSuite))
}

func (s *PreloadSuite) TestPreload() {
	s.Run("returns true if the Preload is valid", func() {
		s.True(validators.Preload("user", []string{"user"}))
	})

	s.Run("returns false if the Preload is invalid", func() {
		s.False(validators.Preload("u", []string{"user"}))
		s.False(validators.Preload("user", []string{"u"}))
		s.False(validators.Preload("users", []string{"user"}))
		s.False(validators.Preload("user", []string{"users"}))
	})
}
