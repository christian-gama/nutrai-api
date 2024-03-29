package mail_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/notify/domain/model"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/go-faker/faker/v4"
)

type MailTestSuite struct {
	suite.Suite
}

func TestMailSuite(t *testing.T) {
	suite.RunUnitTest(t, new(MailTestSuite))
}

func (s *MailTestSuite) TestNewMail() {
	type Sut struct {
		Sut  func() (*mail.Mail, error)
		Data *mail.Mail
	}

	makeSut := func() *Sut {
		data := fake.Mail()

		sut := func() (*mail.Mail, error) {
			return data.Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewMail (Error)", func() {
		s.Run("Context", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()

				sut.Data.Context = nil

				mail, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(mail)
			})
		})

		s.Run("Subject", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()

				sut.Data.Subject = ""

				mail, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(mail)
			})
		})

		s.Run("Template", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()

				sut.Data.Template = nil

				mail, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(mail)
			})
		})

		s.Run("To", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()

				sut.Data.To = []*value.To{}

				mail, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(mail)
			})

			s.Run("Should return an error when contains an empty email", func() {
				sut := makeSut()

				sut.Data.To = []*value.To{
					{Email: "", Name: "John Doe"},
				}

				mail, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(mail)
			})

			s.Run("Should return an error when contains an empty name", func() {
				sut := makeSut()

				sut.Data.To = []*value.To{
					{Email: faker.Email(), Name: ""},
				}

				mail, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(mail)
			})
		})
	})

	s.Run("Attachments", func() {
		s.Run("Should return an error when contains an empty filename", func() {
			sut := makeSut()

			sut.Data.SetAttachments(value.NewAttachment())

			mail, err := sut.Sut()

			s.ErrorAsRequired(err)
			s.Nil(mail)
		})

		s.Run("Should return an error when contains an empty disposition", func() {
			sut := makeSut()

			sut.Data.SetAttachments(value.NewAttachment().
				SetFilename(faker.URL()).
				SetDisposition(""),
			)

			mail, err := sut.Sut()

			s.ErrorAsRequired(err)
			s.Nil(mail)
		})
	})

	s.Run("TestNewMail (Success)", func() {
		s.Run("Should return a mail when all fields are valid", func() {
			sut := makeSut()

			mail, err := sut.Sut()

			s.NoError(err)
			s.NotNil(mail)
			s.Equal(sut.Data.Context, mail.Context, "Context should be equal")
			s.Equal(sut.Data.Subject, mail.Subject, "Subject should be equal")
			s.Equal(sut.Data.Template, mail.Template, "TemplatePath should be equal")
			s.Equal(sut.Data.To, mail.To, "To should be equal")
		})
	})
}
