package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/manager"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/exception/infra/persistence/sql/schema"
	"gorm.io/gorm"
)

// exceptionSQLImpl is the SQL implementation of repo.Exception.
type exceptionSQLImpl struct {
	manager *manager.Manager[exception.Exception, schema.Exception]
}

// NewSQLException returns a new Exception.
func NewSQLException(db *gorm.DB) repo.Exception {
	if db == nil {
		panic(errors.New("db cannot be nil"))
	}

	return &exceptionSQLImpl{
		manager: manager.NewManager[exception.Exception, schema.Exception](db),
	}
}

// DeleteOld implements repo.Exception.
func (p *exceptionSQLImpl) DeleteOld(
	ctx context.Context,
	input repo.DeleteOldExceptionInput,
) error {
	db := p.manager.DB.WithContext(ctx)
	var exceptionSchema schema.Exception

	result := db.
		Model(&exceptionSchema).
		Where("created_at >= ?", input.BeforeDate).
		Delete(&exceptionSchema)
	if result.Error != nil {
		return fmt.Errorf("could not delete exception: %w", result.Error)
	}

	fmt.Println("rows affected:", result.RowsAffected)

	return nil
}

// Save implements repo.Exception.
func (p *exceptionSQLImpl) Save(
	ctx context.Context,
	input repo.SaveExceptionInput,
) (*exception.Exception, error) {
	return p.manager.Save(ctx, manager.SaveInput[exception.Exception]{Model: input.Exception})
}
