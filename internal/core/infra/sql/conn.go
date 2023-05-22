package sql

import (
	"log"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
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
	const maxRetries = 5
	const retryInterval = 1 * time.Second

	c.log.Infof("\tConnecting to SQL database")
	connect := func() (*gorm.DB, error) {
		db, err := gorm.Open(c.dialector(Dsn()), c.opt)
		if err != nil {
			return nil, err
		}
		return db, nil
	}

	ticker := time.NewTicker(retryInterval)
	defer ticker.Stop()

	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = connect()
		if err == nil {
			break
		}
		<-ticker.C
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after %d retries", maxRetries)
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
