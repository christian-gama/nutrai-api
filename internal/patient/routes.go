package patient

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/router/routing"
	"github.com/christian-gama/nutrai-api/internal/patient/api/controller"
)

// Routes registers the routes for user module.
func Routes() *routing.Routing {
	return &routing.Routing{
		Group: "/v1/patients",
		Routes: []*routing.Route{
			{Controller: controller.MakeAllPatients()},
			{Controller: controller.MakeUpdatePatient()},
			{Controller: controller.MakeFindPatient()},
		},
	}
}
