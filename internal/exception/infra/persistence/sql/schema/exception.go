package schema

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/table"
	"github.com/christian-gama/nutrai-api/internal/exception/module"
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
	return table.Name(module.Module, "exceptions")
}
