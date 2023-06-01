package worker

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

// Register creates a number of goroutines to execute the given function.
// If the number of goroutines is greater than the number of CPUs, the number of
// goroutines will be limited to the number of CPUs.
func Register(fn func(), workersAmount int) {
	worker := "workers"
	if workersAmount == 1 {
		worker = "worker"
	}

	if workersAmount > runtime.NumCPU() {
		workersAmount = runtime.NumCPU()
	}

	log.Loading(
		"Creating %s %s %s",
		log.LoadingDetailColor(fmt.Sprint(workersAmount)),
		log.LoadingColor("%s for", worker),
		log.LoadingDetailColor(name(fn)),
	)
	for i := 0; i < workersAmount; i++ {
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
