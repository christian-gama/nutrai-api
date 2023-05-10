package fake

import (
	"fmt"
	"runtime"
)

func ErrGenerating(err error) {
	pc, _, _, _ := runtime.Caller(1)
	callingFunction := runtime.FuncForPC(pc).Name()

	panic(fmt.Errorf("error while generating fake %s: %w", callingFunction, err))
}
