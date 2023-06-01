package value_test

import (
	"fmt"
	"testing"

	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type TemplateTestSuite struct {
	suite.Suite
}

func TestTemplateSuite(t *testing.T) {
	suite.RunUnitTest(t, new(TemplateTestSuite))
}

func (s *TemplateTestSuite) TestNewTemplate() {
	type Sut struct {
		Sut func() *value.Template
	}

	makeSut := func() *Sut {
		sut := func() *value.Template {
			return value.NewTemplate("test.html")
		}

		return &Sut{Sut: sut}
	}

	s.Run("Should return a mail when all fields are valid", func() {
		sut := makeSut()

		mail := sut.Sut()

		s.NotNil(mail)
		s.Equal(
			mail.Path(),
			fmt.Sprintf("%s/%s", mail.BaseDir(), "test.html"),
		)
	})
}
