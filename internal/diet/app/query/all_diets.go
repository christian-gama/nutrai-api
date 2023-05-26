package query

import "github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"

// AllDietInput represents the input data for getting all diets.
type AllDietsInput struct {
	querying.Pagination `faker:"-"`
	querying.Filter     `form:"filter" faker:"-" validate:"query,filter=age weightKG heightM bmi"`
	querying.Sort       `form:"sort" faker:"-" validate:"query,sort=id age weightKG heightM bmi"`
	querying.Preload    `form:"preload" faker:"-" validate:"query,preload=user users"`
}

// AllDietOutput represents the output data for getting all diets.
type AllDietsOutput = FindDietOutput
