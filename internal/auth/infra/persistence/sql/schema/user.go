package schema

import (
	"github.com/christian-gama/nutrai-api/internal/auth/module"
	"github.com/christian-gama/nutrai-api/internal/core/infra/table"
)

// User is the database schema for users.
type User struct {
	ID       uint `gorm:"primaryKey"`
	Email    string
	Name     string
	Password string
}

// TableName returns the table name for the User schema.
func (u User) TableName() string {
	return table.Name(module.Module, "users")
}
