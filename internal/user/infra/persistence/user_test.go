package persistence_test

import (
	"context"
	"testing"

	queryingPort "github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/querying"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/user/infra/persistence"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/user"
	fixture "github.com/christian-gama/nutrai-api/testutils/fixture/user"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"gorm.io/gorm"
)

type UserSuite struct {
	suite.SuiteWithConn
	User func(db *gorm.DB) repo.User
}

func TestUserSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(UserSuite))
}

func (s *UserSuite) SetupTest() {
	s.User = func(db *gorm.DB) repo.User {
		return persistence.NewUser(db)
	}
}

func (s *UserSuite) TestSave() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.SaveUserInput) (*user.User, error)
		Ctx   context.Context
		Input repo.SaveUserInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		user := fake.User()
		input := repo.SaveUserInput{
			User: user,
		}

		sut := s.User(db).Save

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should create a new user", func(db *gorm.DB) {
		sut := makeSut(db)

		user, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.NotZero(user.ID, "Should have an ID")
	})

	s.Run("Should return an error when the user already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		_, err := sut.Sut(sut.Ctx, sut.Input)
		s.NoError(err)

		_, err = sut.Sut(sut.Ctx, sut.Input)
		s.Error(err)
	})
}

func (s *UserSuite) TestDelete() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.DeleteUserInput) error
		Ctx   context.Context
		Input repo.DeleteUserInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.DeleteUserInput{}
		sut := s.User(db).Delete

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should delete a user", func(db *gorm.DB) {
		sut := makeSut(db)

		userDeps := fixture.SaveUser(db, nil)

		sut.Input.IDs = []value.ID{userDeps.User.ID}

		err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
	})

	s.Run("Should delete nothing if the user does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.IDs = []value.ID{404_404_404}

		err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
	})
}

func (s *UserSuite) TestFind() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.FindUserInput,
			preload ...string,
		) (*user.User, error)
		Ctx   context.Context
		Input repo.FindUserInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.FindUserInput{
			ID: 0,
		}
		sut := s.User(db).Find

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should find a user", func(db *gorm.DB) {
		sut := makeSut(db)

		userDeps := fixture.SaveUser(db, nil)

		sut.Input.ID = userDeps.User.ID

		user, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Equal(user.ID, userDeps.User.ID, "Should have the same ID")
	})

	s.Run("Should return an error if the user does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, sut.Input)

		s.Error(err)
	})
}

func (s *UserSuite) TestAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.AllUsersInput,
			preload ...string,
		) (*queryingPort.PaginationOutput[*user.User], error)
		Ctx   context.Context
		Input repo.AllUsersInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.AllUsersInput{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.User(db).All

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should find all users", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.SaveUser(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "Should have a valid id")
		s.Equal(length, result.Total, "Should return %d total", length)
		s.Len(result.Results, length, "Should return %d results", length)
	})

	s.Run("Should return the correct users using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		userDeps := fixture.SaveUser(db, nil)
		length := 3
		for i := 0; i < length; i++ {
			fixture.SaveUser(db, nil)
		}

		sut.Input.Filterer = sut.Input.Filterer.Add(
			"name",
			querying.EqOperator,
			userDeps.User.Name,
		)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Equal(result.Results[0].ID, userDeps.User.ID, "Should have the same ID")
		s.Equal(1, result.Total, "Should return only one user")
		s.Len(result.Results, 1, "Should return only one user")
	})

	s.Run("Should return the correct users using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.SaveUser(db, nil)
		}

		sut.Input.Sorter = sut.Input.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Greater(int(result.Results[1].ID), int(result.Results[2].ID), "Should have the correct order")
	})

	s.Run("Should return the correct users using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.SaveUser(db, nil)
		}

		sut.Input.Sorter = sut.Input.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Greater(int(result.Results[2].ID), int(result.Results[1].ID), "Should have the correct order")
	})

	s.Run("Should return the correct users using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		users := make([]*user.User, 0)
		for i := 0; i < 3; i++ {
			userDeps := fixture.SaveUser(db, nil)
			users = append(users, userDeps.User)
		}

		sut.Input.Paginator = sut.Input.Paginator.SetLimit(1).SetPage(1)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Equal(3, result.Total, "Should return the correct total")
		s.Len(result.Results, 1, "Should return the correct number of users")
		s.Equal(int(users[0].ID), int(result.Results[0].ID), "Should return the correct user")
	})
}

func (s *UserSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.UpdateUserInput,
		) error
		Ctx   context.Context
		Input repo.UpdateUserInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.UpdateUserInput{
			User: fake.User(),
			ID:   1,
		}
		sut := s.User(db).Update

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should update a user", func(db *gorm.DB) {
		sut := makeSut(db)

		userDeps := fixture.SaveUser(db, nil)

		*sut.Input.User = *userDeps.User
		sut.Input.User.Name = "new name"
		sut.Input.ID = userDeps.User.ID

		err := sut.Sut(sut.Ctx, sut.Input)

		s.Require().NoError(err)
		user, err := s.User(db).Find(sut.Ctx, repo.FindUserInput{ID: userDeps.User.ID})
		s.NoError(err)
		s.EqualValues("new name", user.Name, "Should have the new name")
	})

	s.Run("Should return an error if tries to update a non existent user", func(db *gorm.DB) {
		sut := makeSut(db)

		fixture.SaveUser(db, nil)

		sut.Input.User.Name = "new name"
		sut.Input.User.ID = 404_404_404

		err := sut.Sut(sut.Ctx, sut.Input)

		s.Error(err)
	})
}
