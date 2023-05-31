package value

import (
	"os"
	"path"
)

type Template struct {
	path    string
	baseDir string
	ext     string
}

// NewTemplatePath creates a new TemplatePath.
func NewTemplate(fileName string) *Template {
	baseDir := path.Join(os.Getenv("PWD"), "templates")
	ext := ".html"

	fullPath := []string{baseDir}
	fullPath = append(fullPath, fileName)
	path := path.Join(fullPath...)
	path += ext

	return &Template{
		baseDir: baseDir,
		ext:     ext,
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

// Ext returns the ext field.
func (t *Template) Ext() string {
	return t.ext
}
