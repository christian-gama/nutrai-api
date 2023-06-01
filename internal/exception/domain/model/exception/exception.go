package exception

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// Exception is a struct that represents an error.
type Exception struct {
	CreatedAt time.Time `json:"createdAt" faker:"time_now"`
	ID        value.ID  `json:"id" faker:"uint"`
	Message   string    `json:"message" faker:"sentence"`
	Stack     string    `json:"stack" faker:"sentence"`
}

// New creates a new Exception.
func New() *Exception {
	return &Exception{
		CreatedAt: time.Now(),
	}
}

// String implements the fmt.Stringer interface.
func (Exception) String() string {
	return "exception"
}

// Validate validates the Exception fields. It implements the validator interface.
func (e *Exception) Validate() (*Exception, error) {
	return e, nil
}

// SetID sets the id field.
func (e *Exception) SetID(id value.ID) *Exception {
	e.ID = id
	return e
}

// SetMessage sets the message field.
func (e *Exception) SetMessage(message string) *Exception {
	e.Message = message
	return e
}

// SetStack sets the stack field.
func (e *Exception) SetStack(stack string) *Exception {
	e.Stack = stack
	return e
}
