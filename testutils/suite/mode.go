package suite

import (
	"fmt"
	"os"
	"testing"

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
	if mode() == integrationTests || mode() == allTests {
		testify.Run(t, s)
	} else {
		t.SkipNow()
	}
}

// RunUnitTest runs the unit tests if the TEST_MODE is 'unit' or 'all'.
func RunUnitTest(t *testing.T, s testify.TestingSuite) {
	if mode() == unitTests || mode() == allTests {
		testify.Run(t, s)
	} else {
		t.SkipNow()
	}
}
