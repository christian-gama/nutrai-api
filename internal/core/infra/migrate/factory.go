package migrate

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
)

func MakeMigrate() *Migrate {
	db, err := conn.GetPsql().DB()
	if err != nil {
		panic(err)
	}

	return New(db)
}
