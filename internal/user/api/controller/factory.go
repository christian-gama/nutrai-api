package controller

import "github.com/christian-gama/nutrai-api/internal/user/app/command"

func MakeSavePatient() SavePatient {
	return NewSavePatient(command.MakeSavePatientHandler())
}
