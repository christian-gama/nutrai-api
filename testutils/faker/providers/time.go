package providers

import (
	"time"
)

func TimeNow() time.Time {
	return time.Now()
}

func TimeZero() time.Time {
	return time.Time{}
}
