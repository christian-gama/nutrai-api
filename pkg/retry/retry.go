package retry

import (
	"time"
)

func Retry(attempts int, callback func() error) error {
	ticker := time.NewTicker(1 * time.Second)
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
