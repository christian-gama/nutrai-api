package value

import (
	"fmt"
	"path"
)

type Template struct {
	Path    string `json:"value" faker:"len=5"`
	BaseDir string `json:"baseDir" faker:"len=5"`
	Ext     string `json:"ext" faker:"len=2"`
}

// NewTemplatePath creates a new TemplatePath.
func NewTemplate() *Template {
	return &Template{}
}

// SetPath sets the Path field.
func (t *Template) SetPath(elem ...string) *Template {
	if len(elem) == 0 {
		return t
	}

	fullPath := []string{t.BaseDir}
	fullPath = append(fullPath, elem...)
	t.Path = path.Join(fullPath...)
	t.Path += t.Ext

	return t
}

// SetBaseDir sets the BaseDir field.
func (t *Template) SetBaseDir(baseDir string) *Template {
	t.BaseDir = baseDir
	return t
}

// SetExt sets the Ext field.
func (t *Template) SetExt(ext string) *Template {
	if ext == "" {
		return t
	}

	t.Ext = ext
	if ext[0] != '.' {
		t.Ext = fmt.Sprintf(".%s", ext)
	}

	return t
}
