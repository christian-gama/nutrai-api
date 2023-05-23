package retry

import (
	"time"
)

// Retry will retry the callback function until it returns nil or the number of attempts is reached.
func Retry(attempts int, sleep time.Duration, callback func() error) error {
	ticker := time.NewTicker(sleep)
	defer ticker.Stop()

	if attempts <= 0 {
		attempts = 1
	}

	var err error
	for i := 0; i < attempts; i++ {
		err = callback()
		if err == nil {
			break
		}
		<-ticker.C
	}

	return err
}
