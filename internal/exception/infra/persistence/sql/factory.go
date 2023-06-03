package persistence

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/repo"
)

func MakeSQLException() repo.Exception {
	return NewSQLException(conn.GetPsql())
}
