package query

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/plan"
)

// FindPlanInput is the input data of the plan query.
type FindPlanInput struct {
	ID coreValue.ID `uri:"id" faker:"uint" validate:"required,min=1"`

	querying.Preload `form:"preload" validate:"query"`
}

// FindPlanOutput is the output data of the plan query.
type FindPlanOutput struct {
	ID     coreValue.ID `json:"id" faker:"uint"`
	DietID coreValue.ID `json:"dietID" faker:"uint"`
	Text   value.Text   `json:"text" faker:"sentence"`
}
