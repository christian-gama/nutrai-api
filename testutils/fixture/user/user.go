package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/user/infra/persistence"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/user"
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

		user, err := persistence.NewUser(db).
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
