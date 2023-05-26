package query

import persistence "github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql"

func MakeAllDietsHandler() AllDietsHandler {
	return NewAllDietsHandler(persistence.MakeSQLDiet())
}

func MakeFindDietHandler() FindDietHandler {
	return NewFindDietHandler(persistence.MakeSQLDiet())
}
