package command

import persistence "github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql"

func MakeSaveDietHandler() SaveDietHandler {
	return NewSaveDietHandler(persistence.MakeSQLDiet())
}

func MakeSavePlanHandler() SavePlanHandler {
	return NewSavePlanHandler(persistence.MakeSQLPlan(), persistence.MakeSQLDiet())
}

func MakeDeletePlanHandler() DeletePlanHandler {
	return NewDeletePlanHandler(persistence.MakeSQLPlan())
}
