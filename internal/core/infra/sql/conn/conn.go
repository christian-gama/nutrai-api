package conn

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/pkg/retry"
	"gorm.io/gorm"
)

// dialector is a function that returns a GORM dialector.
type dialector func(dsn string) gorm.Dialector

type conn struct {
	dialector
	opt *gorm.Config
	log logger.Logger
}

// NewConn creates a new instance of a GORM connection.
func NewConn(dialector dialector, opt *gorm.Config, logger logger.Logger) *conn {
	return &conn{dialector: dialector, opt: opt, log: logger}
}

// Open will open a new GORM connection.
func (c *conn) Open() (db *gorm.DB) {
	const attempts = 90

	c.log.Loading("\tConnecting to SQL database")

	var err error
	err = retry.Retry(attempts, time.Second, func() error {
		db, err = gorm.Open(c.dialector(dsn()), c.opt)
		return err
	})
	if err != nil {
		c.log.Fatalf("\tFailed to connect to database after %d retries: %v", attempts, err)
	}

	return c.connectionPool(db)
}

// connectionPool will setup the connection pool.
func (c *conn) connectionPool(db *gorm.DB) *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
