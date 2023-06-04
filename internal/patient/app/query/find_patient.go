package query

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
)

// FindPatientInput is the input data of the patient query.
type FindPatientInput struct {
	ID coreValue.ID `uri:"id" faker:"uint" validate:"required,min=1"`

	querying.Preload `form:"preload" faker:"-" validate:"query,preload=allergies"`
}

// FindPatientOutput is the output data of the patient query.
type FindPatientOutput struct {
	ID        coreValue.ID                  `json:"id"`
	Age       value.Age                     `json:"age"`
	HeightM   value.HeightM                 `json:"heightM"`
	WeightKG  value.WeightKG                `json:"weightKG"`
	BMI       value.BMI                     `json:"bmi"`
	Allergies []*FindPatientAllergiesOutput `json:"allergies,omitempty"`
}

type FindPatientAllergiesOutput struct {
	ID        coreValue.ID  `json:"id" faker:"uint"`
	Name      value.Allergy `json:"name" faker:"sentence"`
	PatientID coreValue.ID  `json:"patientID" faker:"uint"`
}
