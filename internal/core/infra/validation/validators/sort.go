package validators

import (
	"strings"

	"github.com/christian-gama/nutrai-api/pkg/slice"
)

// Sort returns true if the string is a valid sort field.
// A valid sort field must have the field name and the order, separated by a space.
// For example: "name:asc" or "name:desc".
func Sort(v string, params []string) bool {
	fieldName := strings.Split(v, ":")[0]
	fieldOrder := strings.Split(v, ":")[1]

	if fieldName == "" {
		return false
	}

	if !slice.Contains(params, fieldName) {
		return false
	}

	if !slice.Contains([]string{"asc", "desc"}, fieldOrder) {
		return false
	}

	return true
}
