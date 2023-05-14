package validators

import (
	"regexp"

	"github.com/christian-gama/nutrai-api/pkg/slice"
)

// Preload returns true if the string is a valid preolad field.
// For example: "preload=users".
func Preload(v string, params []string) bool {
	reg := `preload=(\w+)`

	if !regexp.MustCompile(reg).MatchString(v) {
		return false
	}

	fieldName := regexp.MustCompile(reg).FindStringSubmatch(v)[1]
	if fieldName == "" {
		return false
	}

	if !slice.Contains(params, fieldName) {
		return false
	}

	value := regexp.MustCompile(reg).FindStringSubmatch(v)[3]

	return value != ""
}
