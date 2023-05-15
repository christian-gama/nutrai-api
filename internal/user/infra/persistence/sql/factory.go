package persistence

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
)

func MakeSQLPatient() repo.Patient {
	return NewSQLPatient(sql.MakePostgres())
}

func MakeSQLUser() repo.User {
	return NewSQLUser(sql.MakePostgres())
}
