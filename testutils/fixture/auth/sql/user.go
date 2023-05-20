package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/sql"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	"gorm.io/gorm"
)

type UserDeps struct {
	User *user.User
}

func SaveUser(db *gorm.DB, deps *UserDeps) *UserDeps {
	if deps == nil {
		deps = &UserDeps{}
	}

	user := deps.User
	if user == nil {
		user = fake.User()

		// Ensure that IDs are not set, so that the database can generate them.
		user.ID = 0

		user, err := persistence.NewSQLUser(db).
			Save(context.Background(), repo.SaveUserInput{
				User: user,
			})
		if err != nil {
			panic(fmt.Errorf("could not create user: %w", err))
		}

		deps.User = user
	}

	return deps
}
