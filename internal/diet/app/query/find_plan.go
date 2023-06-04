package query

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
)

// FindPlanInput is the input data of the diet query.
type FindPlanInput struct {
	ID coreValue.ID `uri:"id" faker:"uint" validate:"required,min=1"`

	querying.Preload `form:"preload" validate:"query"`
}

// FindPlanOutput is the output data of the diet query.
type FindPlanOutput struct {
	ID     coreValue.ID    `json:"id"`
	Diet   *FindDietOutput `json:"diet,omitempty"`
	DietID coreValue.ID    `json:"dietID"`
	Text   value.Plan      `json:"text"`
}
