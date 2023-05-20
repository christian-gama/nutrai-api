package query

import persistence "github.com/christian-gama/nutrai-api/internal/patient/infra/persistence/sql"

func MakeAllPatientsHandler() AllPatientsHandler {
	return NewAllPatientsHandler(persistence.MakeSQLPatient())
}

func MakeFindPatientHandler() FindPatientHandler {
	return NewFindPatientHandler(persistence.MakeSQLPatient())
}
