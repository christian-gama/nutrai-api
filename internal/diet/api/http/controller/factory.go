package controller

import (
	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	"github.com/christian-gama/nutrai-api/internal/diet/app/service"
)

func MakeAllPlans() AllPlans {
	return NewAllPlans(query.MakeAllPlansHandler())
}

func MakeFindPlan() FindPlan {
	return NewFindPlan(query.MakeFindPlanHandler())
}

func MakeSavePlan() SavePlan {
	return NewSavePlan(service.MakeSavePlanHandler())
}

func MakeDeletePlan() DeletePlan {
	return NewDeletePlan(command.MakeDeletePlanHandler())
}
