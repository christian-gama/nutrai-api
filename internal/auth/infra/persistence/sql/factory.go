package persistence

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql"
)

func MakeSQLUser() repo.User {
	return NewSQLUser(sql.MakePostgres())
}
