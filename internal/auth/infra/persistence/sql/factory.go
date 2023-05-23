package persistence

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
)

func MakeSQLUser() repo.User {
	return NewSQLUser(conn.MakePostgres())
}
