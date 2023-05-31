package value

import (
	"os"
	"path"
)

type Template struct {
	Path    string `json:"value" faker:"len=5"`
	baseDir string
	ext     string
}

// NewTemplatePath creates a new TemplatePath.
func NewTemplate() *Template {
	return &Template{
		baseDir: path.Join(os.Getenv("PWD"), "templates"),
		ext:     ".html",
	}
}

// SetPath sets the Path field.
func (t *Template) SetPath(fileName string) *Template {
	fullPath := []string{t.baseDir}
	fullPath = append(fullPath, fileName)
	t.Path = path.Join(fullPath...)
	t.Path += t.ext

	return t
}

// BaseDir returns the baseDir field.
func (t *Template) BaseDir() string {
	return t.baseDir
}

// Ext returns the ext field.
func (t *Template) Ext() string {
	return t.ext
}
