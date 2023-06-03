package gpt

import (
	value "github.com/christian-gama/nutrai-api/internal/gpt/domain/value/model"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

type Model struct {
	Name             value.Name             `faker:"sentence"`
	MaxTokens        value.MaxTokens        `faker:"boundary_start=1024, boundary_end=1024"`
	Temperature      value.Temperature      `faker:"boundary_start=0.0, boundary_end=1.0"` // 0.0 - 1.0
	TopP             value.TopP             `faker:"boundary_start=0.0, boundary_end=1.0"` // 0.0 - 1.0
	N                value.N                `faker:"boundary_start=1, boundary_end=100"`   // number of responses
	Stop             []value.Stop           `faker:"-"`
	PresencePenalty  value.PresencePenalty  `faker:"boundary_start=-2.0, boundary_end=2.0"` // -2.0 to 2.0 - Number beetwen -2.0 and 2.0
	FrequencyPenalty value.FrequencyPenalty `faker:"boundary_start=-2.0, boundary_end=2.0"` // -2.0 to 2.0 - Number beetwen -2.0 and 2.0
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

	if err := m.Temperature.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := m.TopP.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := m.N.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := m.PresencePenalty.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := m.FrequencyPenalty.Validate(); err != nil {
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

// SetTemperature sets the Temperature of the Model.
func (m *Model) SetTemperature(temperature value.Temperature) *Model {
	m.Temperature = temperature
	return m
}

// SetTopP sets the TopP of the Model.
func (m *Model) SetTopP(topP value.TopP) *Model {
	m.TopP = topP
	return m
}

// SetN sets the N of the Model.
func (m *Model) SetN(n value.N) *Model {
	m.N = n
	return m
}

// SetStop sets the Stop of the Model.
func (m *Model) SetStop(stop []value.Stop) *Model {
	m.Stop = stop
	return m
}

// SetPresencePenalty sets the PresencePenalty of the Model.
func (m *Model) SetPresencePenalty(presencePenalty value.PresencePenalty) *Model {
	m.PresencePenalty = presencePenalty
	return m
}

// SetFrequencyPenalty sets the FrequencyPenalty of the Model.
func (m *Model) SetFrequencyPenalty(frequencyPenalty value.FrequencyPenalty) *Model {
	m.FrequencyPenalty = frequencyPenalty
	return m
}
