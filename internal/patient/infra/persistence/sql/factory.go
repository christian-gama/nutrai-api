package persistence

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
)

func MakeSQLPatient() repo.Patient {
	return NewSQLPatient(sql.MakePostgres())
}
