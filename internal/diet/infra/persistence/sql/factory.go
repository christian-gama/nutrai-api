package persistence

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
)

func MakeSQLDiet() repo.Diet {
	return NewSQLDiet(conn.MakePostgres())
}
