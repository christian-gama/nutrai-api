package controller

import (
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/christian-gama/nutrai-api/internal/user/app/query"
)

func MakeAllPatients() AllPatients {
	return NewAllPatients(query.MakeAllPatientsHandler())
}

func MakeSavePatient() SavePatient {
	return NewSavePatient(command.MakeSavePatientHandler())
}

func MakeUpdatePatient() UpdatePatient {
	return NewUpdatePatient(command.MakeUpdatePatientHandler())
}
