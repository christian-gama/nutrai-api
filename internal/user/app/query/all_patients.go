package query

import "github.com/christian-gama/nutrai-api/internal/shared/infra/querying"

type AllPatientsInput struct {
	querying.Filter `form:"filter"`
	querying.Pagination
	querying.Sort `form:"sort"`
}

type AllPatientsOutput = FindPatientOutput
