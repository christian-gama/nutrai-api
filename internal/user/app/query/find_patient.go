package query

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"
)

// FindPatientInput is the input data of the patient query.
type FindPatientInput struct {
	ID coreValue.ID `url:"id" faker:"uint"`

	querying.Preload `form:"preload" faker:"-"`
}

// FindPatientOutput is the output data of the patient query.
type FindPatientOutput struct {
	ID       coreValue.ID    `json:"id"`
	Age      value.Age       `json:"age"`
	HeightM  value.HeightM   `json:"heightM"`
	WeightKG value.WeightKG  `json:"weightKG"`
	User     *FindUserOutput `json:"user,omitempty"`
	BMI      value.BMI       `json:"bmi"`
}
