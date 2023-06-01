package conn

import (
	"time"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/pkg/retry"
	"gorm.io/gorm"
)

// dialector is a function that returns a GORM dialector.
type dialector func(dsn string) gorm.Dialector

// NewConn creates a new instance of a GORM connection.
func NewConn(dialector dialector, opt *gorm.Config) (db *gorm.DB) {
	const attempts = 90

	log.Loading("\tConnecting to SQL database")

	var err error
	err = retry.Retry(attempts, time.Second, func() error {
		db, err = gorm.Open(dialector(dsn()), opt)
		return err
	})
	if err != nil {
		log.Fatalf("\tFailed to connect to database after %d retries: %v", attempts, err)
	}

	return connectionPool(db)
}

// connectionPool will setup the connection pool.
func connectionPool(db *gorm.DB) *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(env.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(env.DB.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(env.DB.ConnMaxLifetime)

	return db
}
