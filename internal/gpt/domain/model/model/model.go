package gpt

import (
	value "github.com/christian-gama/nutrai-api/internal/gpt/domain/value/model"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

type Model struct {
	Name      value.Name      `faker:"sentence"`
	MaxTokens value.MaxTokens `faker:"boundary_start=1024, boundary_end=1024"`
}

// NewModel returns a new Model instance.
func NewModel() *Model {
	return &Model{}
}

// String implements the fmt.Stringer interface.
func (Model) String() string {
	return "model"
}

// Validate returns an error if the model is invalid.
func (m *Model) Validate() (*Model, error) {
	var errs *errutil.Error

	if err := m.Name.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := m.MaxTokens.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if errs.HasErrors() {
		return nil, errs
	}

	return m, nil
}

// SetName sets the Name of the Model.
func (m *Model) SetName(name value.Name) *Model {
	m.Name = name
	return m
}

// SetMaxTokens sets the MaxTokens of the Model.
func (m *Model) SetMaxTokens(maxTokens value.MaxTokens) *Model {
	m.MaxTokens = maxTokens
	return m
}
