package controller_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type PathSuite struct {
	suite.Suite
}

func TestPathSuite(t *testing.T) {
	suite.RunUnitTest(t, new(PathSuite))
}

func (s *PathSuite) TestAddPath() {
	s.Run("returns the joined path", func() {
		path := controller.JoinPath("resource", "id")

		s.EqualValues("/resource/id", path)
	})

	s.Run("add slash if path is empty", func() {
		path := controller.JoinPath("", "id")

		s.EqualValues("/id", path)
	})

	s.Run("panic if uses slash on the path name", func() {
		s.Panics(func() {
			controller.JoinPath("resource/", "id")
		})
	})
}
