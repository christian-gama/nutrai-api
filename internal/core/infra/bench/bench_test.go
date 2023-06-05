package bench_test

import (
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/infra/bench"
	"github.com/christian-gama/nutrai-api/testutils/suite"
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
		s.WithinDuration(expected, actual, 100*time.Millisecond)
	})
}

var wait = 1 * time.Microsecond

func sampleFunc() {
	time.Sleep(wait)
}
