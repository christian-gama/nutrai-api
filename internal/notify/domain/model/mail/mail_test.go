package mail_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/notify/domain/model"
	"github.com/christian-gama/nutrai-api/testutils/suite"
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
			return mail.NewMail().
				SetBody(data.Body).
				SetFrom(data.From).
				SetSubject(data.Subject).
				SetTemplate(data.Template).
				SetTo(data.To).
				Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewMail (Error)", func() {
		s.Run("Body", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()

				sut.Data.Body = ""

				mail, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(mail)
			})
		})

		s.Run("From", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()

				sut.Data.From = ""

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

				sut.Data.Template = ""

				mail, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(mail)
			})
		})

		s.Run("To", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()

				sut.Data.To = []value.Email{}

				mail, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(mail)
			})

			s.Run("Should return an error when contains an empty email", func() {
				sut := makeSut()

				sut.Data.To = []value.Email{""}

				mail, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(mail)
			})
		})
	})

	s.Run("TestNewMail (Success)", func() {
		s.Run("Should return a mail when all fields are valid", func() {
			sut := makeSut()

			mail, err := sut.Sut()

			s.NoError(err)
			s.NotNil(mail)
			s.Equal(sut.Data.Body, mail.Body, "Body should be equal")
			s.Equal(sut.Data.From, mail.From, "From should be equal")
			s.Equal(sut.Data.Subject, mail.Subject, "Subject should be equal")
			s.Equal(sut.Data.Template, mail.Template, "Template should be equal")
			s.Equal(sut.Data.To, mail.To, "To should be equal")
		})
	})
}
