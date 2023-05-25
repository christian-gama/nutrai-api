package routes

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router/routes"
	"github.com/christian-gama/nutrai-api/internal/patient/api/http/controller"
)

// Register registers the routes for this module.
func Register() {
	routes.
		Api("/v1/patients").
		SetController(controller.MakeAllPatients()).
		SetController(controller.MakeSavePatient()).
		SetController(controller.MakeUpdatePatient()).
		SetController(controller.MakeFindPatient())
}
