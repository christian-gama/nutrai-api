package query

import (
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql/querying"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"
)

// FindPatientInput is the input data of the patient query.
type FindPatientInput struct {
	ID sharedvalue.ID `form:"id" faker:"uint"`

	querying.Preload `form:"preload" faker:"-"`
}

// FindPatientOutput is the output data of the patient query.
type FindPatientOutput struct {
	ID       sharedvalue.ID  `json:"id"`
	Age      value.Age       `json:"age"`
	HeightM  value.HeightM   `json:"heightM"`
	WeightKG value.WeightKG  `json:"weightKG"`
	User     *FindUserOutput `json:"user,omitempty"`
	BMI      value.BMI       `json:"bmi"`
}
