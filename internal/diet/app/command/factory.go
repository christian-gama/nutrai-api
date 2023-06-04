package command

import persistence "github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql"

func MakeSaveDietHandler() SaveDietHandler {
	return NewSaveDietHandler(persistence.MakeSQLDiet())
}

func MakeDeletePlanHandler() DeletePlanHandler {
	return NewDeletePlanHandler(persistence.MakeSQLPlan())
}
