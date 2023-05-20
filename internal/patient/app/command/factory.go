package command

import (
	persistence "github.com/christian-gama/nutrai-api/internal/patient/infra/persistence/sql"
)

func MakeSavePatientHandler() SavePatientHandler {
	return NewSavePatientHandler(persistence.MakeSQLPatient())
}

func MakeUpdatePatientHandler() UpdatePatientHandler {
	return NewUpdatePatientHandler(
		persistence.MakeSQLPatient(),
	)
}
