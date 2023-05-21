package asserts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// HasChanged asserts that the given schema has changed.
func HasChanged(
	t *testing.T,
	oldObj, newObj any,
	msgAndArgs ...any,
) bool {
	if len(msgAndArgs) > 0 {
		return assert.NotEqual(t, oldObj, newObj, msgAndArgs...)
	}

	return assert.NotEqual(
		t,
		oldObj,
		newObj,
		"The schema should have changed, but it did not.",
	)
}

// HasNotChanged asserts that the given schema has not changed.
func HasNotChanged(
	t *testing.T,
	oldObj, newObj any,
	msgAndArgs ...any,
) bool {
	if len(msgAndArgs) > 0 {
		return assert.Equal(t, oldObj, newObj, msgAndArgs...)
	}

	return assert.Equal(
		t,
		oldObj,
		newObj,
		"The should not have changed, but it did.",
	)
}
