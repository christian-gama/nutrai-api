package query

import "github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"

// AllPlanInput represents the input data for getting all plans.
type AllPlansInput struct {
	querying.Pagination `faker:"-"`
	querying.Filter     `form:"filter" faker:"-" validate:"query,filter=dietID text"`
	querying.Sort       `form:"sort" faker:"-" validate:"query,sort=id text"`
	querying.Preload    `form:"preload" faker:"-" validate:"query,preload=diet"`
}

// AllPlanOutput represents the output data for getting all plans.
type AllPlansOutput = FindPlanOutput
