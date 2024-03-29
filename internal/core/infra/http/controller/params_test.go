package controller_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type ParamsSuite struct {
	suite.Suite
}

func TestParamsSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ParamsSuite))
}

func (s *ParamsSuite) TestAddParams() {
	s.Run("returns a Params with the given param", func() {
		params := controller.AddParams("id")

		s.EqualValues(controller.Params{"id"}, params)
	})

	s.Run("returns a Params with the given params", func() {
		params := controller.AddParams("id").Add("name")

		s.EqualValues(controller.Params{"id", "name"}, params)
	})

	s.Run("returns a path with multiple params", func() {
		params := controller.AddParams("id").Add("name")

		s.EqualValues("/resource/:id/:name", params.ToPath("/resource"))
	})

	s.Run("returns a path with one param", func() {
		params := controller.AddParams("id")

		s.EqualValues("/resource/:id", params.ToPath("/resource"))
	})
}
