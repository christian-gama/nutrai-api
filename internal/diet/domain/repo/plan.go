package repo

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
)

// SavePlanInput is the input for the Save method.
type SavePlanInput struct {
	Plan *plan.Plan
}

// AllPlansInput is the input for the All method.
type AllPlansInput struct {
	queryer.Filterer
	queryer.Sorter
	queryer.Paginator
	queryer.Preloader
}

// FindPlanInput is the input for the Find method.
type FindPlanInput struct {
	ID value.ID
	queryer.Filterer
	queryer.Preloader
}

// DeletePlanInput is the input for the Delete method.
type DeletePlanInput struct {
	IDs []value.ID
}

// UpdatePlanInput is the input for the Update method.
type UpdatePlanInput struct {
	Plan *plan.Plan
	ID   value.ID
}

// Plan is the interface that wraps the basic Plan methods.
type Plan interface {
	// All returns all plans.
	All(input AllPlansInput) (*queryer.PaginationOutput[*plan.Plan], error)

	// Delete deletes the plan with the given id.
	Delete(input DeletePlanInput) error

	// Find returns the plan with the given id.
	Find(input FindPlanInput) (*plan.Plan, error)

	// Save saves the given plan.
	Save(input SavePlanInput) (*plan.Plan, error)

	// Update updates the given plan.
	Update(input UpdatePlanInput) error
}
