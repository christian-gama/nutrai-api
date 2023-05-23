package migrate

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
)

func MakeMigrate() *Migrate {
	db, err := conn.MakePostgres().DB()
	if err != nil {
		panic(err)
	}

	return New(db, log.MakeWithCaller())
}
