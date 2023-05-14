package querying_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"gorm.io/gorm"
)

type PreloadSuite struct {
	suite.SuiteWithSQLConn
}

func TestPreloadSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(PreloadSuite))
}

func (s *PreloadSuite) TestPreload() {
	type YetAnotherEntity struct {
		gorm.Model
		Field string `gorm:"not null"`
	}

	type AnotherEntity struct {
		gorm.Model
		Field              string            `gorm:"not null"`
		YetAnotherEntityID uint              `gorm:"not null"`
		YetAnotherEntity   *YetAnotherEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}

	type Entity struct {
		gorm.Model
		Field           string         `gorm:"not null"`
		AnotherEntityID uint           `gorm:"not null"`
		AnotherEntity   *AnotherEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}

	type Sut func(*gorm.DB, queryer.Preloader) *Entity

	makeSut := func(tx *gorm.DB) Sut {
		return func(tx *gorm.DB, Preloader queryer.Preloader) *Entity {
			tx.AutoMigrate(&Entity{})

			entity := &Entity{
				Field: "a",
				AnotherEntity: &AnotherEntity{
					Field: "b",
					YetAnotherEntity: &YetAnotherEntity{
						Field: "c",
					},
				},
			}
			if err := tx.Create(&entity).Error; err != nil {
				s.FailNow(err.Error())
			}

			if err := tx.
				Scopes(querying.PreloadScope(Preloader)).
				First(&entity).Error; err != nil {
				s.FailNow(err.Error())
			}

			tx.Migrator().DropTable(&Entity{})

			return entity
		}
	}

	s.Run("no panic if Preload is nil", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entity := sut(tx, nil)

		s.Equal("a", entity.Field)
	})

	s.Run("no panic if Preload is empty string", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entity := sut(tx, querying.Preload{""})

		s.Equal("a", entity.Field)
	})

	s.Run("preloads nested relations", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entity := sut(tx, querying.Preload{
			"AnotherEntity.YetAnotherEntity",
		})

		s.Equal("a", entity.Field)
		s.Equal("b", entity.AnotherEntity.Field)
		s.Equal("c", entity.AnotherEntity.YetAnotherEntity.Field)
	})

	s.Run("preloads the correct nested relations", func(tx *gorm.DB) {
		sut := makeSut(tx)

		entity := sut(tx, querying.Preload{
			"AnotherEntity",
		})

		s.Equal("a", entity.Field)
		s.Equal("b", entity.AnotherEntity.Field)
		s.Nil(entity.AnotherEntity.YetAnotherEntity)
	})

	s.Run("preloads using different string cases", func(tx *gorm.DB) {
		s.Run("lower case", func(tx *gorm.DB) {
			sut := makeSut(tx)

			entity := sut(tx, querying.Preload{
				"anotherEntity.yetAnotherEntity",
			})

			s.Equal("a", entity.Field)
			s.Equal("b", entity.AnotherEntity.Field)
			s.Equal("c", entity.AnotherEntity.YetAnotherEntity.Field)
		})

		s.Run("title case", func(tx *gorm.DB) {
			sut := makeSut(tx)

			entity := sut(tx, querying.Preload{
				"AnotherEntity.YetAnotherEntity",
			})

			s.Equal("a", entity.Field)
			s.Equal("b", entity.AnotherEntity.Field)
			s.Equal("c", entity.AnotherEntity.YetAnotherEntity.Field)
		})

		s.Run("snake case", func(tx *gorm.DB) {
			sut := makeSut(tx)

			entity := sut(tx, querying.Preload{
				"another_entity.yet_another_entity",
			})

			s.Equal("a", entity.Field)
			s.Equal("b", entity.AnotherEntity.Field)
			s.Equal("c", entity.AnotherEntity.YetAnotherEntity.Field)
		})

		s.Run("mixed case", func(tx *gorm.DB) {
			sut := makeSut(tx)

			entity := sut(tx, querying.Preload{
				"AnotherEntity.yet_anotherEntity",
			})

			s.Equal("a", entity.Field)
			s.Equal("b", entity.AnotherEntity.Field)
			s.Equal("c", entity.AnotherEntity.YetAnotherEntity.Field)
		})
	})
}
