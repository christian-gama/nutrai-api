package value_test

import (
	"errors"
	"testing"

	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type AttachmentTestSuite struct {
	suite.Suite
}

func TestAttachmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(AttachmentTestSuite))
}

func (s *AttachmentTestSuite) TestAttachment() {
	type Sut struct {
		Sut *value.Attachment
	}

	makeSut := func() *Sut {
		return &Sut{Sut: value.NewAttachment()}
	}

	s.Run("TestAttachment (Success)", func() {
		sut := makeSut()

		attachment := sut.Sut.SetFilename("test.png").
			SetDisposition("inline").
			SetContentID("test").
			SetContentType(value.ContentTypePNG).
			SetName("test")

		s.Equal(attachment.Filename, "test.png")
		s.Equal(attachment.Disposition, "inline")
		s.Equal(attachment.ContentID(), "test")
		s.Equal(attachment.ContentType(), "image/png")
		s.Equal(attachment.Content(func(name string) ([]byte, error) {
			return []byte("test"), nil
		}), "dGVzdA==")
	})

	s.Run("TestAttachment (Error)", func() {
		s.Run("Content", func() {
			s.Run("Should panic when fileReader returns an error", func() {
				sut := makeSut()

				attachment := sut.Sut.SetFilename("test.png").
					SetDisposition("inline").
					SetContentID("test")

				s.Panics(func() {
					attachment.Content(func(name string) ([]byte, error) {
						return nil, errors.New("error")
					})
				})
			})
		})

		s.Run("ContentType", func() {
			s.Run("Should automatically set the content type based on the filename", func() {
				sut := makeSut()

				files := []string{
					"test.invalid",
					"test.jpegg",
				}
				for _, file := range files {
					sut.Sut.SetFilename(file)

					s.Panics(func() {
						sut.Sut.ContentType()
					})
				}
			})
		})
	})

	s.Run("Should automatically set the content type based on the filename", func() {
		sut := makeSut()

		files := []string{
			"test.png",
			"test.jpeg",
			"test.jpg",
			"test.txt",
			"test.pdf",
			"test.csv",
		}
		for _, file := range files {
			sut.Sut.SetFilename(file)

			s.NotPanics(func() {
				sut.Sut.ContentType()
			})
		}
	})

	s.Run("Should automatically set the content id based on the filename", func() {
		sut := makeSut()

		sut.Sut.SetFilename("/this/is/a/path/test.png")

		s.Equal(sut.Sut.ContentID(), "test")
	})

	s.Run("Should automatically set the name based on the filename", func() {
		sut := makeSut()

		sut.Sut.SetFilename("/this/is/a/path/test.png")

		s.Equal(sut.Sut.Name(), "test")
	})
}
