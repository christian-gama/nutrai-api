package manager_test

import (
	"context"
	"testing"

	queryingPort "github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/manager"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

type ManagerSuite struct {
	suite.SuiteWithConn
	Sample func(db *gorm.DB) manager.Repository[Sample]
}

type Sample struct {
	ID   value.ID `gorm:"primaryKey" faker:"-"`
	Name string   `gorm:"not null,unique"`
}

func (s Sample) TableName() string {
	return "samples"
}

func FakeSample() *Sample {
	data := new(Sample)
	err := faker.FakeData(data)
	if err != nil {
		panic(err)
	}

	return data
}

func SaveSample(db *gorm.DB) *Sample {
	sample := FakeSample()
	err := db.Create(sample).Error
	if err != nil {
		panic(err)
	}

	return sample
}

func TestManagerSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(ManagerSuite))
}

func (s *ManagerSuite) SetupTest() {
	s.Sample = func(db *gorm.DB) manager.Repository[Sample] {
		return manager.NewManager[Sample, Sample](db)
	}

	db := sql.MakePostgres()
	db.AutoMigrate(&Sample{})
}

func (s *ManagerSuite) AfterTest() {
	db := sql.MakePostgres()
	db.Migrator().DropTable(&Sample{})
}

func (s *ManagerSuite) TestSave() {
	type Sut struct {
		Sut   func(ctx context.Context, input manager.SaveInput[Sample]) (*Sample, error)
		Ctx   context.Context
		Input manager.SaveInput[Sample]
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		sample := FakeSample()
		input := manager.SaveInput[Sample]{
			Model: sample,
		}

		sut := s.Sample(db).Save

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should create a new sample", func(db *gorm.DB) {
		sut := makeSut(db)

		sample, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.NotZero(sample.ID, "Should have an ID")
	})

	s.Run("Should return an error when the sample already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		_, err := sut.Sut(sut.Ctx, sut.Input)
		s.NoError(err)

		_, err = sut.Sut(sut.Ctx, sut.Input)
		s.Error(err)
	})
}

func (s *ManagerSuite) TestDelete() {
	type Sut struct {
		Sut   func(ctx context.Context, input manager.DeleteInput[Sample]) error
		Ctx   context.Context
		Input manager.DeleteInput[Sample]
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := manager.DeleteInput[Sample]{}
		sut := s.Sample(db).Delete

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should delete a sample", func(db *gorm.DB) {
		sut := makeSut(db)

		sample := SaveSample(db)

		sut.Input.IDs = []value.ID{sample.ID}

		err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
	})

	s.Run("Should delete nothing if the sample does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.IDs = []value.ID{404_404_404}

		err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
	})
}

func (s *ManagerSuite) TestFind() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input manager.FindInput[Sample],
		) (*Sample, error)
		Ctx   context.Context
		Input manager.FindInput[Sample]
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := manager.FindInput[Sample]{
			ID: 0,
		}
		sut := s.Sample(db).Find

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should find a sample", func(db *gorm.DB) {
		sut := makeSut(db)

		sampleFixture := SaveSample(db)

		sut.Input.ID = sampleFixture.ID

		sample, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Equal(sample.ID, sampleFixture.ID)
	})

	s.Run("Should return an error if the sample does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, sut.Input)

		s.Error(err)
	})
}

func (s *ManagerSuite) TestAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input manager.AllInput[Sample],
		) (*queryingPort.PaginationOutput[*Sample], error)
		Ctx   context.Context
		Input manager.AllInput[Sample]
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := manager.AllInput[Sample]{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
			Preloader: querying.Preload{},
		}
		sut := s.Sample(db).All

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should find all samples", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			SaveSample(db)
		}

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "Should have a valid id")
		s.Equal(length, result.Total, "Should return %d total", length)
		s.Len(result.Results, length, "Should return %d results", length)
	})

	s.Run("Should return the correct samples using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		sample := SaveSample(db)
		length := 3
		for i := 0; i < length; i++ {
			SaveSample(db)
		}

		sut.Input.Filterer = sut.Input.Filterer.Add(
			"name",
			querying.EqOperator,
			sample.Name,
		)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Equal(result.Results[0].ID, sample.ID, "Should have the same ID")
		s.Equal(1, result.Total, "Should return only one sample")
		s.Len(result.Results, 1, "Should return only one sample")
	})

	s.Run("Should return the correct samples using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			SaveSample(db)
		}

		sut.Input.Sorter = sut.Input.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Greater(int(result.Results[1].ID), int(result.Results[2].ID), "Should have the correct order")
	})

	s.Run("Should return the correct samples using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			SaveSample(db)
		}

		sut.Input.Sorter = sut.Input.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Greater(int(result.Results[2].ID), int(result.Results[1].ID), "Should have the correct order")
	})

	s.Run("Should return the correct samples using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		samples := make([]*Sample, 0)
		for i := 0; i < 3; i++ {
			sample := SaveSample(db)
			samples = append(samples, sample)
		}

		sut.Input.Paginator = sut.Input.Paginator.SetLimit(1).SetPage(1)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Equal(3, result.Total, "Should return the correct total")
		s.Len(result.Results, 1, "Should return the correct number of samples")
		s.Equal(int(samples[0].ID), int(result.Results[0].ID), "Should return the correct sample")
	})
}

func (s *ManagerSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input manager.UpdateInput[Sample],
		) error
		Ctx   context.Context
		Input manager.UpdateInput[Sample]
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := manager.UpdateInput[Sample]{
			Model: FakeSample(),
			ID:    1,
		}
		sut := s.Sample(db).Update

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should update a sample", func(db *gorm.DB) {
		sut := makeSut(db)

		sample := SaveSample(db)

		*sut.Input.Model = *sample
		sut.Input.Model.Name = "new name"
		sut.Input.ID = sample.ID

		err := sut.Sut(sut.Ctx, sut.Input)

		s.Require().NoError(err)
		sample, err = s.Sample(db).Find(sut.Ctx, manager.FindInput[Sample]{ID: sample.ID})
		s.NoError(err)
		s.EqualValues("new name", sample.Name, "Should have the new name")
	})
}
