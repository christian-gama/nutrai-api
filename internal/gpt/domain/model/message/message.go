package gpt

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/gpt/domain/value/message"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

type Message struct {
	ID      coreValue.UUID `faker:"uuid_digit"`
	Role    value.Role     `faker:"-"`
	Content value.Content  `faker:"sequence"`
	Tokens  value.Tokens   `faker:"boundary_start=1, boundary_end=1024"`
}

// NewMessage returns a new Message instance.
func NewMessage() *Message {
	return &Message{}
}

// String implements the fmt.Stringer interface.
func (Message) String() string {
	return "message"
}

// Validate returns an error if the message is invalid.
func (m *Message) Validate() (*Message, error) {
	var errs *errutil.Error

	if err := m.Role.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := m.Content.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := m.Tokens.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if errs.HasErrors() {
		return nil, errs
	}

	return m, nil
}
