package querying_test

import (
	"testing"
	"time"

	queryingPort "github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/querying"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"gorm.io/gorm"
)

type SortSuite struct {
	suite.SuiteWithConn
}

func TestSortSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(SortSuite))
}

func (s *SortSuite) TestSort() {
	type Entity struct {
		gorm.Model
		Name string `gorm:"not null"`
	}

	type Sut func(*gorm.DB, queryingPort.Sorter) []*Entity

	makeSut := func(tx *gorm.DB) Sut {
		return func(tx *gorm.DB, sorter queryingPort.Sorter) []*Entity {
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
				Scopes(querying.SortScope(sorter)).
				Find(&entities).
				Error; err != nil {
				s.FailNow(err.Error())
			}

			tx.Migrator().DropTable(&Entity{})

			return entities
		}
	}

	s.Run("sorts by name asc", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.AddSort("name", false))

		s.Len(entities, 5)
		s.Equal("a", entities[0].Name)
		s.Equal("b", entities[1].Name)
		s.Equal("c", entities[2].Name)
		s.Equal("d", entities[3].Name)
		s.Equal("d", entities[4].Name)
	})

	s.Run("sorts by name desc", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.AddSort("name", true))

		s.Len(entities, 5)
		s.Equal("d", entities[0].Name)
		s.Equal("d", entities[1].Name)
		s.Equal("c", entities[2].Name)
		s.Equal("b", entities[3].Name)
		s.Equal("a", entities[4].Name)
	})

	s.Run("sorts by id desc", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.AddSort("id", false))

		s.Len(entities, 5)
		s.Equal(1, int(entities[0].ID))
		s.Equal(2, int(entities[1].ID))
		s.Equal(3, int(entities[2].ID))
		s.Equal(4, int(entities[3].ID))
		s.Equal(5, int(entities[4].ID))
	})

	s.Run("sorts by id desc", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.AddSort("id", true))

		s.Len(entities, 5)
		s.Equal(5, int(entities[0].ID))
		s.Equal(4, int(entities[1].ID))
		s.Equal(3, int(entities[2].ID))
		s.Equal(2, int(entities[3].ID))
		s.Equal(1, int(entities[4].ID))
	})

	s.Run("sorts by created_at desc", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.AddSort("created_at", false))

		s.Len(entities, 5)
		s.Equal(1, int(entities[0].ID))
		s.Equal(2, int(entities[1].ID))
		s.Equal(3, int(entities[2].ID))
		s.Equal(4, int(entities[3].ID))
		s.Equal(5, int(entities[4].ID))
	})

	s.Run("sorts by created_at desc", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.AddSort("created_at", true))

		s.Len(entities, 5)
		s.Equal(5, int(entities[0].ID))
		s.Equal(4, int(entities[1].ID))
		s.Equal(3, int(entities[2].ID))
		s.Equal(2, int(entities[3].ID))
		s.Equal(1, int(entities[4].ID))
	})

	s.Run("sorts by name asc, id desc", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.
			AddSort("name", false).
			Add("id", true))

		s.Len(entities, 5)
		s.Equal("a", entities[0].Name)
		s.Equal(1, int(entities[0].ID))
		s.Equal("b", entities[1].Name)
		s.Equal(2, int(entities[1].ID))
		s.Equal("c", entities[2].Name)
		s.Equal(3, int(entities[2].ID))
		s.Equal("d", entities[3].Name)
		s.Equal(5, int(entities[3].ID))
		s.Equal("d", entities[4].Name)
		s.Equal(4, int(entities[4].ID))
	})

	s.Run("sorts by name desc, id asc", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.
			AddSort("name", true).
			Add("id", false))

		s.Len(entities, 5)
		s.Equal("d", entities[0].Name)
		s.Equal(4, int(entities[0].ID))
		s.Equal("d", entities[1].Name)
		s.Equal(5, int(entities[1].ID))
		s.Equal("c", entities[2].Name)
		s.Equal(3, int(entities[2].ID))
		s.Equal("b", entities[3].Name)
		s.Equal(2, int(entities[3].ID))
		s.Equal("a", entities[4].Name)
		s.Equal(1, int(entities[4].ID))
	})

	s.Run("no panic if sort is empty string", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entities := sut(tx, querying.Sort{""})

		s.Len(entities, 5)
	})
}
