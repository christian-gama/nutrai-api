package user

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/router/routing"
	"github.com/christian-gama/nutrai-api/internal/user/api/controller"
)

// Routes registers the routes for user module.
func Routes() *routing.Routing {
	return &routing.Routing{
		Group: "/v1/user",
		Routes: []*routing.Route{
			{Controller: controller.MakeAllPatients()},
			{Controller: controller.MakeSavePatient()},
			{Controller: controller.MakeUpdatePatient()},
			{Controller: controller.MakeFindPatient()},
		},
	}
}
