package query

import "github.com/christian-gama/nutrai-api/internal/shared/infra/querying"

// AllPatientsInput represents the input data for getting all patients.
type AllPatientsInput struct {
	querying.Filter `form:"filter"`
	querying.Pagination
	querying.Sort `form:"sort"`
}

// AllPatientsOutput represents the output data for getting all patients.
type AllPatientsOutput = FindPatientOutput
