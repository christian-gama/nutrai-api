package bench_test

import (
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/infra/bench"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/logger"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/mock"
)

type BenchSuite struct {
	suite.Suite
}

func TestBenchSuite(t *testing.T) {
	suite.RunUnitTest(t, new(BenchSuite))
}

func (s *BenchSuite) TestDuration() {
	s.Run("returns the duration of the function", func() {
		duration := bench.Duration(sampleFunc)

		expected := time.Time{}.Add(wait)
		actual := time.Time{}.Add(duration)

		s.NotEqual(expected, actual)
		s.WithinDuration(expected, actual, 10*time.Millisecond)
	})
}

func (s *BenchSuite) TestPrintDuration() {
	s.Run("prints the duration of the function", func() {
		log := mocks.NewLogger(s.T())
		resource := "test"
		log.On("Infof", "%s took %dms to complete", resource, mock.AnythingOfType("int64")).
			Return().
			Once()

		bench.PrintDuration(log, resource, sampleFunc)

		log.AssertNumberOfCalls(s.T(), "Infof", 1)
		log.AssertCalled(
			s.T(),
			"Infof",
			"%s took %dms to complete",
			resource,
			mock.AnythingOfType("int64"),
		)
	})
}

var wait = 1 * time.Microsecond

func sampleFunc() {
	time.Sleep(wait)
}
