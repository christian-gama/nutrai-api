package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/exception/infra/persistence/sql"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/exception/domain/model/exception"
	"gorm.io/gorm"
)

type ExceptionDeps struct {
	Exception *exception.Exception
}

func SaveException(db *gorm.DB, deps *ExceptionDeps) *ExceptionDeps {
	if deps == nil {
		deps = &ExceptionDeps{}
	}

	exception := deps.Exception
	if exception == nil {
		exception = fake.Exception()

		exception, err := persistence.NewSQLException(db).
			Save(context.Background(), repo.SaveExceptionInput{
				Exception: exception,
			})
		if err != nil {
			panic(fmt.Errorf("could not create exception: %w", err))
		}

		deps.Exception = exception
	}

	return deps
}
