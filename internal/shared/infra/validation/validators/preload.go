package validators

import (
	"github.com/christian-gama/nutrai-api/pkg/slice"
)

// Preload returns true if the string is a valid preolad field.
// For example: "user".
func Preload(v string, params []string) bool {
	return slice.Contains(params, v)
}
