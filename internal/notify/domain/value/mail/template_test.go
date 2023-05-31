package value_test

import (
	"fmt"
	"testing"

	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/go-faker/faker/v4"
)

type TemplateTestSuite struct {
	suite.Suite
}

func TestTemplateSuite(t *testing.T) {
	suite.RunUnitTest(t, new(TemplateTestSuite))
}

func (s *TemplateTestSuite) TestNewTemplate() {
	type Sut struct {
		Sut  func() *value.Template
		Data *value.Template
	}

	makeSut := func() *Sut {
		data := new(value.Template)
		err := faker.FakeData(data)
		if err != nil {
			fake.ErrGenerating(err)
		}

		sut := func() *value.Template {
			return value.NewTemplate().SetPath(data.Path)
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("Should return a mail when all fields are valid", func() {
		sut := makeSut()

		mail := sut.Sut()

		s.NotNil(mail)
		s.Equal(
			mail.Path,
			fmt.Sprintf("%s/%s%s", mail.BaseDir(), sut.Data.Path, mail.Ext()),
		)
	})
}
