package value

import "path"

type TemplatePath struct {
	Value string `json:"value" faker:"len=50"`
}

// NewTemplatePath creates a new TemplatePath.
func NewTemplatePath(paths ...string) *TemplatePath {
	return &TemplatePath{Value: path.Join(paths...)}
}

func (t *TemplatePath) Path() string {
	return t.Value
}
