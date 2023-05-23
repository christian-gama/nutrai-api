package controller

import (
	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
)

func MakeAllPatients() AllPatients {
	return NewAllPatients(query.MakeAllPatientsHandler())
}

func MakeUpdatePatient() UpdatePatient {
	return NewUpdatePatient(command.MakeUpdatePatientHandler())
}

func MakeFindPatient() FindPatient {
	return NewFindPatient(query.MakeFindPatientHandler())
}

func MakeSavePatient() SavePatient {
	return NewSavePatient(command.MakeSavePatientHandler())
}
