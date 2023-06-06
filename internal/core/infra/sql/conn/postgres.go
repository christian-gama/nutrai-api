package conn

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"gorm.io/gorm"
)

var psql *conn

// GetPsql returns the postgres connection.
func GetPsql() *gorm.DB {
	if psql == nil {
		log.Fatal(
			errors.InternalServerError(
				"postgres connection does not exist - did you forget to initialize it?",
			),
		)
	}

	return psql.DB
}
