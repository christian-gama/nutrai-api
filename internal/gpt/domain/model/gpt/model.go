package gpt

import (
	"github.com/christian-gama/nutrai-api/config/env"
	value "github.com/christian-gama/nutrai-api/internal/gpt/domain/value/gpt"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
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
	return &Model{
		Name:             value.Name(env.Gpt.Model),
		MaxTokens:        value.MaxTokens(env.Gpt.MaxTokens),
		Temperature:      value.Temperature(env.Gpt.Temperature),
		TopP:             value.TopP(env.Gpt.TopP),
		N:                value.N(env.Gpt.N),
		PresencePenalty:  value.PresencePenalty(env.Gpt.PresencePenalty),
		FrequencyPenalty: value.FrequencyPenalty(env.Gpt.FrequencyPenalty),
	}
}

// String implements the fmt.Stringer interface.
func (Model) String() string {
	return "model"
}

// Validate returns an error if the model is invalid.
func (m *Model) Validate() (*Model, error) {
	var errs *errutil.Error

	if m.Name == "" {
		errs = errutil.Append(errs, errors.Required("Name"))
	}

	if m.MaxTokens <= 0 {
		errs = errutil.Append(errs, errors.Invalid("MaxTokens", "MaxTokens must be greater than 0"))
	}

	if m.Temperature < 0.0 || m.Temperature > 1.0 {
		errs = errutil.Append(
			errs,
			errors.Invalid("Temperature", "Temperature must be between 0.0 and 1.0"),
		)
	}

	if m.TopP < 0.0 || m.TopP > 1.0 {
		errs = errutil.Append(errs, errors.Invalid("TopP", "TopP must be between 0.0 and 1.0"))
	}

	if m.N <= 0 {
		errs = errutil.Append(errs, errors.Invalid("N", "N must be greater than 0"))
	}

	if m.PresencePenalty < -2.0 || m.PresencePenalty > 2.0 {
		errs = errutil.Append(
			errs,
			errors.Invalid("PresencePenalty", "PresencePenalty must be between -2.0 and 2.0"),
		)
	}

	if m.FrequencyPenalty < -2.0 || m.FrequencyPenalty > 2.0 {
		errs = errutil.Append(
			errs,
			errors.Invalid("FrequencyPenalty", "FrequencyPenalty must be between -2.0 and 2.0"),
		)
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
