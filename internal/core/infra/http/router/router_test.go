package router_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type RouterSuite struct {
	suite.Suite
}

func TestRouterSuite(t *testing.T) {
	suite.RunUnitTest(t, new(RouterSuite))
}

func (s *RouterSuite) TestRouter() {
	s.Todo("Should have recovery wrapping the whole router, catching panics", func() {})
}
