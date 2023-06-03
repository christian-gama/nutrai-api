package conn

import (
	"errors"

	"gorm.io/gorm"
)

var psql *conn

// GetPsql returns the postgres connection.
func GetPsql() *gorm.DB {
	if psql == nil {
		panic(errors.New("postgres connection does not exist"))
	}

	return psql.DB
}

// ClosePsql closes the postgres connection.
func ClosePsql() {
	if psql == nil {
		panic(errors.New("postgres connection does not exist"))
	}

	db, err := psql.DB.DB()
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
