package querying_test

import (
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"gorm.io/gorm"
)

type FilterSuite struct {
	suite.SuiteWithSQLConn
}

func TestFilterSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(FilterSuite))
}

func (s *FilterSuite) TestFilter() {
	type Entity struct {
		gorm.Model
		Name string `gorm:"not null"`
	}

	type Sut func(*gorm.DB, queryer.Filterer) []*Entity

	makeSut := func(tx *gorm.DB) Sut {
		return func(tx *gorm.DB, filter queryer.Filterer) []*Entity {
			tx.AutoMigrate(&Entity{})

			if err := tx.CreateInBatches([]*Entity{
				{Name: "abc", Model: gorm.Model{ID: 1, CreatedAt: time.Now()}},
				{Name: "bcd", Model: gorm.Model{ID: 2, CreatedAt: time.Now().Add(time.Second)}},
				{Name: "CdE", Model: gorm.Model{ID: 3, CreatedAt: time.Now().Add(time.Second * 2)}},
				{Name: "def", Model: gorm.Model{ID: 4, CreatedAt: time.Now().Add(time.Second * 3)}},
				{Name: "ghi", Model: gorm.Model{ID: 5, CreatedAt: time.Now().Add(time.Second * 4)}},
			}, 5).
				Error; err != nil {
				s.FailNow(err.Error())
			}

			entities := []*Entity{}

			if err := tx.
				Scopes(querying.FilterScope(filter)).
				Find(&entities).
				Error; err != nil {
				s.FailNow(err.Error())
			}

			tx.Migrator().DropTable(&Entity{})

			return entities
		}
	}

	s.Run("filters by name", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.AddFilter("name", querying.EqOperator, "bcd"))

		s.Equal(1, len(entities))
		s.Equal("bcd", entities[0].Name)
	})

	s.Run("returns no results if no match", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.
			AddFilter("name", querying.EqOperator, "bcd").
			Add("id", querying.EqOperator, "1"))

		s.Equal(0, len(entities))
	})

	s.Run("returns all results if no filter is provided", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.Filter{})

		s.Equal(5, len(entities))
	})

	s.Run("returns a result based on an ilike query", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.AddFilter("name", querying.LikeOperator, "bc"))

		s.Equal(2, len(entities))
		s.Equal("abc", entities[0].Name)
		s.Equal("bcd", entities[1].Name)
	})

	s.Run("returns a result based on an ilike query with wildcards", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.AddFilter("name", querying.LikeOperator, "%bc%"))

		s.Equal(2, len(entities))
		s.Equal("abc", entities[0].Name)
		s.Equal("bcd", entities[1].Name)
	})

	s.Run("returns a result based on a IN query", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.AddFilter("name", querying.InOperator, []string{"abc", "bcd"}))

		s.Equal(2, len(entities))
		s.Equal("abc", entities[0].Name)
		s.Equal("bcd", entities[1].Name)
	})

	s.Run("do not panic if filter is empty string", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.Filter{""})

		s.Equal(5, len(entities))
	})
}
