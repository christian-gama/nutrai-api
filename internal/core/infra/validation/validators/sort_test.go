package validators_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/validation/validators"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type SortSuite struct {
	suite.Suite
}

func TestSortSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SortSuite))
}

func (s *SortSuite) TestSort() {
	s.Run("returns true if the sort is valid", func() {
		s.True(validators.Sort("name:asc", []string{"name"}))
		s.True(validators.Sort("name:desc", []string{"name"}))
		s.True(validators.Sort("name:asc", []string{"name", "id"}))
		s.True(validators.Sort("name:desc", []string{"name", "id"}))
	})

	s.Run("returns false if the sort is invalid", func() {
		s.False(validators.Sort("name:asc", []string{"id"}))
		s.False(validators.Sort("name:desc", []string{"id"}))
		s.False(validators.Sort("name:invalid", []string{"name"}))
	})
}
