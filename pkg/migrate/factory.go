package migrate

import "github.com/christian-gama/nutrai-api/internal/shared/infra/sql"

func MakeMigrate(silent bool) *Migrate {
	db, err := sql.MakePostgres().DB()
	if err != nil {
		panic(err)
	}

	return New(db)
}
