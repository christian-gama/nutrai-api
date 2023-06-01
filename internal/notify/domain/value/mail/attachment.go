package value

import (
	"encoding/base64"
	"path"
	"strings"

	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// Attachment encapsulates information regarding an email attachment.
// This struct is used for handling attached files or data within email messages.
type Attachment struct {
	Filename    string
	Disposition string
	contentType string
	contentID   string
	content     string
	name        string
}

// NewAttachment creates a new Attachment.
func NewAttachment() *Attachment {
	return &Attachment{}
}

// SetContentID sets the ContentID field.
func (a *Attachment) SetContentID(contentID string) *Attachment {
	a.contentID = contentID
	return a
}

// SetFilename sets the Filename field.
func (a *Attachment) SetFilename(filename string) *Attachment {
	a.Filename = filename
	return a
}

// SetType sets the Type field.
func (a *Attachment) SetContentType(contentType string) *Attachment {
	a.contentType = contentType
	return a
}

// SetContent sets the Content field.
func (a *Attachment) Content(fileReader func(name string) ([]byte, error)) string {
	if a.content != "" {
		return a.content
	}

	fileBytes, err := fileReader(a.Filename)
	if err != nil {
		panic(errors.InternalServerError("could not read file: %s", err.Error()))
	}

	a.content = base64.StdEncoding.EncodeToString(fileBytes)

	return a.content
}

// SetName sets the Name field.
func (a *Attachment) SetName(name string) *Attachment {
	a.name = name
	return a
}

// SetDisposition sets the Disposition field.
func (a *Attachment) SetDisposition(disposition string) *Attachment {
	a.Disposition = disposition
	return a
}

// ContentID returns the ContentID field. If the content ID is empty, it returns the Name field.
func (a *Attachment) ContentID() string {
	if a.contentID == "" {
		return a.Name()
	}

	return a.contentID
}

// Name returns the Name field. If the name is empty, it returns the base filename of the
// attachment without the extension.
func (a *Attachment) Name() string {
	if a.name == "" {
		ext := path.Ext(a.Filename)
		return strings.TrimSuffix(path.Base(a.Filename), ext)
	}

	return a.name
}

// ContentType returns the ContentType field. If the content type is empty, it returns the
// content type based on the file extension.
func (a *Attachment) ContentType() string {
	if a.contentType == "" {
		return a.contentTypeFromFilename()
	}

	return a.contentType
}

func (a *Attachment) contentTypeFromFilename() string {
	mapContentType := map[string]string{
		".png":  ContentTypePNG,
		".jpg":  ContentTypeJPG,
		".gif":  ContentTypeGIF,
		".jpeg": ContentTypeJPEG,
		".pdf":  ContentTypePDF,
		".zip":  ContentTypeZIP,
		".csv":  ContentTypeCSV,
		".txt":  ContentTypeTXT,
	}

	ext := path.Ext(a.Filename)
	contentType, ok := mapContentType[ext]
	if !ok {
		panic(errors.InternalServerError("could not find content type for extension: %s", ext))
	}

	return contentType
}

const (
	ContentTypePNG  = "image/png"
	ContentTypeJPG  = "image/jpg"
	ContentTypeGIF  = "image/gif"
	ContentTypeJPEG = "image/jpeg"
	ContentTypePDF  = "application/pdf"
	ContentTypeZIP  = "application/zip"
	ContentTypeCSV  = "text/csv"
	ContentTypeTXT  = "text/plain"
)
