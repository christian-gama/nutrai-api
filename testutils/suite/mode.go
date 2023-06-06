package suite

import (
	"fmt"
	"os"
	"testing"

	redisconn "github.com/christian-gama/nutrai-api/internal/core/infra/redis/conn"
	sqlconn "github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	testify "github.com/stretchr/testify/suite"
)

const (
	UnitTests        = "unit"
	IntegrationTests = "integration"
	AllTests         = "all"
)

// Mode returns the test Mode to be executed.
func Mode() string {
	modes := []string{UnitTests, IntegrationTests, AllTests}
	mode, ok := os.LookupEnv("TEST_MODE")
	if !ok {
		return AllTests
	}

	if !slice.Contains(modes, mode) {
		panic(fmt.Errorf("expected TEST_MODE to be one of: %v", modes))
	}

	return mode
}

// RunIntegrationTest runs the integration tests if the TEST_MODE is 'integration' or all.
func RunIntegrationTest(t *testing.T, s testify.TestingSuite) {
	t.Parallel()
	if Mode() == IntegrationTests || Mode() == AllTests {
		sqlconn.MakePsql()
		redisconn.MakeRedis()

		testify.Run(t, s)
	} else {
		t.SkipNow()
	}
}

// RunIntegrationTestOnce runs the integration tests if the TEST_MODE is 'integration' or all and
// sets the ran flag to true, avoiding running the test twice.
func RunIntegrationTestOnce(t *testing.T, s testify.TestingSuite, ran *bool) {
	t.Parallel()
	if Mode() == IntegrationTests || Mode() == AllTests {
		if !*ran {
			*ran = true
			sqlconn.MakePsql()
			redisconn.MakeRedis()

			testify.Run(t, s)
		} else {
			t.SkipNow()
		}
	}
}

// RunUnitTest runs the unit tests if the TEST_MODE is 'unit' or 'all'.
func RunUnitTest(t *testing.T, s testify.TestingSuite) {
	t.Parallel()
	if Mode() == UnitTests || Mode() == AllTests {
		testify.Run(t, s)
	} else {
		t.SkipNow()
	}
}

// RunUnitTestOnce runs the unit tests if the TEST_MODE is 'unit' or 'all' and sets the ran flag to
// true, avoiding running the test twice.
func RunUnitTestOnce(t *testing.T, s testify.TestingSuite, ran *bool) {
	t.Parallel()
	if !*ran {
		if Mode() == UnitTests || Mode() == AllTests {
			*ran = true
			testify.Run(t, s)
		} else {
			t.SkipNow()
		}
	}
}
