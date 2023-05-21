package persistence

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/repo"
)

func MakeSQLException() repo.Exception {
	return NewSQLException(sql.MakePostgres())
}
