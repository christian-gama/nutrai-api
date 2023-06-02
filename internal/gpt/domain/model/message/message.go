package gpt

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	gpt "github.com/christian-gama/nutrai-api/internal/gpt/domain/model/model"
	value "github.com/christian-gama/nutrai-api/internal/gpt/domain/value/message"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

type Message struct {
	ID      coreValue.UUID `faker:"uuid_digit"`
	Role    value.Role     `faker:"-"`
	Content value.Content  `faker:"sequence"`
	Tokens  value.Tokens   `faker:"boundary_start=1, boundary_end=1024"`
	Model   *gpt.Model     `faker:"-"`
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

// SetModel sets the Model of the Message.
func (m *Message) SetModel(model *gpt.Model) *Message {
	m.Model = model
	return m
}
