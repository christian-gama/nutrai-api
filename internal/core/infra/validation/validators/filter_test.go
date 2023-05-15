package validators_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/validation/validators"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type FilterSuite struct {
	suite.Suite
}

func TestFilterSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FilterSuite))
}

func (s *FilterSuite) TestFilter() {
	s.Run("returns true if the filter is valid", func() {
		s.True(validators.Filter("field=name,op=eq,value=John", []string{"name"}))
		s.True(validators.Filter("field=name,op=like,value=John", []string{"name"}))
		s.True(validators.Filter("field=name,op=in,value=[1 2 3]", []string{"name"}))
		s.True(validators.Filter("field=name,op=gt,value=1", []string{"name"}))
		s.True(validators.Filter("field=name,op=gte,value=1", []string{"name"}))
		s.True(validators.Filter("field=name,op=lt,value=1", []string{"name"}))
		s.True(validators.Filter("field=name,op=lte,value=1", []string{"name"}))
		s.True(validators.Filter("field=name,op=neq,value=1", []string{"name"}))
		s.True(validators.Filter("field=name,op=eq,value=John", []string{"name", "id"}))
		s.True(validators.Filter("field=name,op=like,value=John", []string{"name", "id"}))
		s.True(validators.Filter("field=name,op=in,value=[1 2 3]", []string{"name", "id"}))
		s.True(validators.Filter("field=name,op=gt,value=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,op=gte,value=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,op=lt,value=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,op=lte,value=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,op=neq,value=1", []string{"name", "id"}))
	})

	s.Run("returns false if the filter is invalid", func() {
		s.False(validators.Filter("field=name,op=eq,value=John", []string{"id"}))
		s.False(validators.Filter("field=name,op=like,value=John", []string{"id"}))
		s.False(validators.Filter("field=name,op=in,value=[1 2 3]", []string{"id"}))
		s.False(validators.Filter("field=name,op=gt,value=1", []string{"id"}))
		s.False(validators.Filter("field=name,op=gte,value=1", []string{"id"}))
		s.False(validators.Filter("field=name,op=lt,value=1", []string{"id"}))
		s.False(validators.Filter("field=name,op=lte,value=1", []string{"id"}))
		s.False(validators.Filter("field=name,op=neq,value=1", []string{"id"}))
		s.False(validators.Filter("field=name,invalid=1", []string{"name"}))
		s.False(validators.Filter("field=name,op=eq,value=John", []string{}))
		s.False(validators.Filter("field=name,op=in,value=John", []string{}))
	})
}
