package http

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	"github.com/christian-gama/nutrai-api/internal/patient/api/http/controller"
)

// Routes registers the routes for user module.
func Routes() {
	router.Routes = append(router.Routes, &router.Routing{
		Group: "/v1/patients",
		Routes: []*router.Route{
			{Controller: controller.MakeAllPatients()},
			{Controller: controller.MakeUpdatePatient()},
			{Controller: controller.MakeFindPatient()},
		},
	})
}
