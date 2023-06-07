package internal

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/auth"
	authMiddleware "github.com/christian-gama/nutrai-api/internal/auth/api/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core"
	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	routesMiddleware "github.com/christian-gama/nutrai-api/internal/core/infra/http/router/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	redisconn "github.com/christian-gama/nutrai-api/internal/core/infra/redis/conn"
	sqlconn "github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
	"github.com/christian-gama/nutrai-api/internal/diet"
	"github.com/christian-gama/nutrai-api/internal/exception"
	expectionMiddleware "github.com/christian-gama/nutrai-api/internal/exception/api/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/metrics"
	metricsMiddleware "github.com/christian-gama/nutrai-api/internal/metrics/api/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/notify"
	"github.com/christian-gama/nutrai-api/internal/patient"
)

// setupModules is responsible for setting up the modules of the application. When creating a new
// module, it should be added here. The order of the modules is important, as some modules may
// behave differently depending on the order.
func setupModules() {
	module.Init(core.Init)
	module.Init(exception.Init)
	module.Init(auth.Init)
	module.Init(metrics.Init)
	module.Init(patient.Init)
	module.Init(diet.Init)
	module.Init(notify.Init)
}

// setupStrategies is responsible for setting up the strategies of the application. When creating a
// new strategy, it should be added here.
func setupStrategies() {
	controller.AuthJwtStrategy.SetMiddleware(authMiddleware.MakeAuthJwt())
	controller.AuthApiKeyStrategy.SetMiddleware(authMiddleware.MakeAuthApiKey())
	routesMiddleware.RecoveryAndPersistStrategy.SetMiddleware(expectionMiddleware.MakeRecovery())
	routesMiddleware.MetricsStrategy.SetMiddleware(metricsMiddleware.MakeMetrics())
}

// setupConnections is responsible for setting up the connections of the application.
func setupConnections() {
	sqlconn.MakePsql()
	redisconn.MakeRedis()
}

// Bootstrap is responsible for booting up the application.
func Bootstrap(envFile string) {
	env.NewLoader(envFile).Load()
	log.SugaredLogger = log.New()
	setupConnections()
	setupStrategies()
	router.SetupRouter()
	setupModules()
}
