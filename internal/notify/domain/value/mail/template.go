package value

import (
	"os"
	"path"

	"github.com/christian-gama/nutrai-api/config/env"
)

type Template struct {
	path    string
	baseDir string
}

// NewTemplatePath creates a new TemplatePath.
func NewTemplate(fileName string) *Template {
	baseDir := path.Join(os.Getenv("PWD"), env.Mailer.TemplatePath)

	fullPath := []string{baseDir}
	fullPath = append(fullPath, fileName)
	path := path.Join(fullPath...)

	return &Template{
		baseDir: baseDir,
		path:    path,
	}
}

// SetPath sets the Path field.
func (t *Template) Path() string {
	return t.path
}

// BaseDir returns the baseDir field.
func (t *Template) BaseDir() string {
	return t.baseDir
}
