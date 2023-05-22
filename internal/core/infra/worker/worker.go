package worker

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

// Create creates a number of goroutines to execute the given function.
// If the number of goroutines is greater than the number of CPUs, the number of
// goroutines will be limited to the number of CPUs.
func Create(fn func(), n int) {
	worker := "workers"
	if n == 1 {
		worker = "worker"
	}

	if n > runtime.NumCPU() {
		n = runtime.NumCPU()
	}

	log.MakeWithCaller().Infof("Creating %d %s for %s", n, worker, name(fn))
	for i := 0; i < n; i++ {
		go fn()
	}
}

// name returns the name of the function.
func name(fn func()) string {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	pathParts := strings.Split(name, "/")
	name = pathParts[len(pathParts)-1]

	nameParts := strings.Split(name, ".")
	name = nameParts[0] + "." + nameParts[1]

	return name
}
