package gpt

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/gpt/domain/value/gpt"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

const (
	User      value.Role = "user"
	System    value.Role = "system"
	Assistant value.Role = "assistant"
)

type Message struct {
	ID      coreValue.UUID `faker:"uuid_digit"`
	Role    value.Role     `faker:"-"`
	Content value.Content  `faker:"sentence"`
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

	if m.Role != User && m.Role != System && m.Role != Assistant {
		errs = errutil.Append(
			errs,
			errors.Invalid("Role", "Role must be one of: user, system, assistant"),
		)
	}

	if m.Content == "" {
		errs = errutil.Append(errs, errors.Required("Content"))
	}

	if m.Tokens <= 0 {
		errs = errutil.Append(errs, errors.Invalid("Tokens", "Tokens must be greater than 0"))
	}

	if errs.HasErrors() {
		return nil, errs
	}

	return m, nil
}

// SetID sets the ID of the Message.
func (m *Message) SetID(id coreValue.UUID) *Message {
	m.ID = id
	return m
}

// SetRole sets the Role of the Message.
func (m *Message) SetRole(role value.Role) *Message {
	m.Role = role
	return m
}

// SetContent sets the Content of the Message.
func (m *Message) SetContent(content value.Content) *Message {
	m.Content = content
	return m
}

// SetTokens sets the Tokens of the Message.
func (m *Message) SetTokens(tokens value.Tokens) *Message {
	m.Tokens = tokens
	return m
}
