package service

import persistence "github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql"

func MakeSavePlanHandler() SavePlanHandler {
	return NewSavePlanHandler(persistence.MakeSQLPlan(), persistence.MakeSQLDiet())
}
