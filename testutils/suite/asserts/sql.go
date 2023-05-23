package asserts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// SQLCount asserts that the given schema has the expected count.
func SQLCount(
	t *testing.T,
	db *gorm.DB,
	schema schema.Tabler,
	expectedCount int,
	msgAndArgs ...any,
) bool {
	var count int64
	db.Model(&schema).Count(&count)

	if len(msgAndArgs) > 0 {
		return assert.EqualValues(t, expectedCount, int(count), msgAndArgs...)
	}

	return assert.EqualValues(
		t,
		expectedCount,
		int(count),
		"Expected the schema to have %d records, but it has %d.",
		expectedCount,
		count,
	)
}

// SQLRecordExist asserts that the given schema exists.
func SQLRecordExist(t *testing.T, db *gorm.DB, schema schema.Tabler, msgAndArgs ...any) bool {
	if len(msgAndArgs) > 0 {
		return assert.NoError(t, db.First(&schema).Error, msgAndArgs...)
	}

	return assert.NoError(t, db.First(&schema).Error, "The record should exist, but it does not.")
}

// SQLRecordDoesNotExist asserts that the given schema does not exist.
func SQLRecordDoesNotExist(
	t *testing.T,
	db *gorm.DB,
	schema schema.Tabler,
	msgAndArgs ...any,
) bool {
	if len(msgAndArgs) > 0 {
		return assert.Error(t, db.First(&schema).Error, msgAndArgs...)
	}

	return assert.Error(t, db.First(&schema).Error, "The record should not exist, but it does.")
}
