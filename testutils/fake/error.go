package fake

import "fmt"

func ErrGenerating(name string, err error) {
	panic(fmt.Errorf("error while generating fake %s: %w", name, err))
}
