package schema

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// Exception is a repository schema.
type Exception struct {
	CreatedAt time.Time
	ID        value.ID `gorm:"primaryKey"`
	Message   string
	Stack     string
}

// TableName returns the table name.
func (Exception) TableName() string {
	return "exceptions"
}
