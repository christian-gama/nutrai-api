package querying_test

import (
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"gorm.io/gorm"
)

type PaginationSuite struct {
	suite.SuiteWithSQLConn
}

func TestPaginationSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(PaginationSuite))
}

func (s *PaginationSuite) TestPagination() {
	type Entity struct {
		gorm.Model
		Name string `gorm:"not null"`
	}

	type Sut func(*gorm.DB, queryer.Paginator) []*Entity

	makeSut := func(tx *gorm.DB) Sut {
		return func(tx *gorm.DB, pagination queryer.Paginator) []*Entity {
			tx.AutoMigrate(&Entity{})

			if err := tx.CreateInBatches([]*Entity{
				{Name: "a", Model: gorm.Model{ID: 1, CreatedAt: time.Now()}},
				{Name: "b", Model: gorm.Model{ID: 2, CreatedAt: time.Now().Add(time.Second)}},
				{Name: "c", Model: gorm.Model{ID: 3, CreatedAt: time.Now().Add(time.Second * 2)}},
				{Name: "d", Model: gorm.Model{ID: 4, CreatedAt: time.Now().Add(time.Second * 3)}},
				{Name: "d", Model: gorm.Model{ID: 5, CreatedAt: time.Now().Add(time.Second * 4)}},
			}, 5).
				Error; err != nil {
				s.FailNow(err.Error())
			}

			entities := []*Entity{}

			if err := tx.
				Scopes(querying.PaginationScope(pagination)).
				Find(&entities).
				Error; err != nil {
				s.FailNow(err.Error())
			}

			tx.Migrator().DropTable(&Entity{})

			return entities
		}
	}

	s.Run("returns only the first page with 1 item", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, &querying.Pagination{
			Page:  1,
			Limit: 1,
		})

		s.Equal(1, len(entities))
		s.Equal(1, int(entities[0].ID))
	})

	s.Run("returns only the second page with 1 item", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, &querying.Pagination{
			Page:  2,
			Limit: 1,
		})

		s.Equal(1, len(entities))
		s.Equal(2, int(entities[0].ID))
	})

	s.Run("returns only the second page with 2 items", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, &querying.Pagination{
			Page:  2,
			Limit: 2,
		})

		s.Equal(2, len(entities))
		s.Equal(3, int(entities[0].ID))
		s.Equal(4, int(entities[1].ID))
	})

	s.Run("returns only the first page with 5 items", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, &querying.Pagination{
			Page:  1,
			Limit: 5,
		})

		s.Equal(5, len(entities))
		s.Equal(1, int(entities[0].ID))
		s.Equal(2, int(entities[1].ID))
		s.Equal(3, int(entities[2].ID))
		s.Equal(4, int(entities[3].ID))
		s.Equal(5, int(entities[4].ID))
	})

	s.Run("returns nothing when page is out of bounds", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, &querying.Pagination{
			Page:  404,
			Limit: 5,
		})

		s.Equal(0, len(entities))
	})
}
