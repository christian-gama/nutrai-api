package query

import "github.com/christian-gama/nutrai-api/internal/shared/infra/sql/querying"

// AllPatientsInput represents the input data for getting all patients.
type AllPatientsInput struct {
	querying.Pagination `faker:"-"`
	querying.Filter     `form:"filter" faker:"-"`
	querying.Sort       `form:"sort" faker:"-"`
	querying.Preload    `form:"preload" faker:"-"`
}

// AllPatientsOutput represents the output data for getting all patients.
type AllPatientsOutput = FindPatientOutput
