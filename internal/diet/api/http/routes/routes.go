package routes

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router/routes"
	"github.com/christian-gama/nutrai-api/internal/diet/api/http/controller"
)

// Register registers the routes for this module.
func Register() {
	routes.
		Api("/v1/diets").
		SetController(controller.MakeAllPlans()).
		SetController(controller.MakeFindPlan()).
		SetController(controller.MakeSavePlan()).
		SetController(controller.MakeDeletePlan())
}
