package migrate

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql"
)

func MakeMigrate() *Migrate {
	db, err := sql.MakePostgres().DB()
	if err != nil {
		panic(err)
	}

	return New(db)
}
