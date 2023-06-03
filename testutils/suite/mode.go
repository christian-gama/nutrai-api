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
	unitTests        = "unit"
	integrationTests = "integration"
	allTests         = "all"
)

// mode returns the test mode to be executed.
func mode() string {
	modes := []string{unitTests, integrationTests, allTests}
	mode, ok := os.LookupEnv("TEST_MODE")
	if !ok {
		return allTests
	}

	if !slice.Contains(modes, mode) {
		panic(fmt.Errorf("expected TEST_MODE to be one of: %v", modes))
	}

	return mode
}

// RunIntegrationTest runs the integration tests if the TEST_MODE is 'integration' or all.
func RunIntegrationTest(t *testing.T, s testify.TestingSuite) {
	sqlconn.MakePsql()
	redisconn.MakeRedis()

	t.Parallel()
	if mode() == integrationTests || mode() == allTests {
		testify.Run(t, s)
	}
}

// RunIntegrationTestOnce runs the integration tests if the TEST_MODE is 'integration' or all and
// sets the ran flag to true, avoiding running the test twice.
func RunIntegrationTestOnce(t *testing.T, s testify.TestingSuite, ran *bool) {
	if !*ran {
		sqlconn.MakePsql()
		redisconn.MakeRedis()

		t.Parallel()
		if mode() == integrationTests || mode() == allTests {
			testify.Run(t, s)
		}

		*ran = true
	}
}

// RunUnitTest runs the unit tests if the TEST_MODE is 'unit' or 'all'.
func RunUnitTest(t *testing.T, s testify.TestingSuite) {
	t.Parallel()
	if mode() == unitTests || mode() == allTests {
		testify.Run(t, s)
	}
}

// RunUnitTestOnce runs the unit tests if the TEST_MODE is 'unit' or 'all' and sets the ran flag to
// true, avoiding running the test twice.
func RunUnitTestOnce(t *testing.T, s testify.TestingSuite, ran *bool) {
	if !*ran {
		t.Parallel()
		if mode() == unitTests || mode() == allTests {
			testify.Run(t, s)
		}

		*ran = true
	}
}
