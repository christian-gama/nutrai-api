package retry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/pkg/retry"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type RetrySuite struct {
	suite.Suite
}

func TestRetrySuite(t *testing.T) {
	suite.RunUnitTest(t, new(RetrySuite))
}

func (s *RetrySuite) TestRetry() {
	sleep := 1 * time.Microsecond

	s.Run("should return nil if callback returns nil", func() {
		err := retry.Retry(1, sleep, func() error {
			return nil
		})

		s.NoError(err)
	})

	s.Run("should return error if callback returns error", func() {
		err := retry.Retry(1, sleep, func() error {
			return errors.New("error")
		})

		s.Error(err)
	})

	s.Run("should retry 10 times if callback returns error", func() {
		attempts := 10
		err := retry.Retry(attempts, sleep, func() error {
			attempts--
			return errors.New("error")
		})

		s.Error(err)
		s.Equal(0, attempts)
	})

	s.Run("should retry 5 times and return nil if callback returns nil", func() {
		attempts := 10
		err := retry.Retry(attempts, sleep, func() error {
			attempts--
			if attempts > 5 {
				return errors.New("error")
			}
			return nil
		})

		s.NoError(err)
		s.Equal(5, attempts)
	})

	s.Run("should retry at least 1 time if attempts is 0", func() {
		attempts := 0
		err := retry.Retry(attempts, sleep, func() error {
			attempts--
			return errors.New("error")
		})

		s.Error(err)
		s.Equal(-1, attempts)
	})
}
