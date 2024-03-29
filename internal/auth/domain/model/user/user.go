package user

import (
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// User represents a User model, containing essential credentials and personal identification
// information for an individual user. This includes a unique identifier, email address, password,
// and name. This model is used to manage user authentication, authorization, and maintain
// user-specific information. This model can be used to validate a user's credentials and identify a
// user across the application,.
type User struct {
	ID       coreValue.ID   `faker:"uint"`
	Email    value.Email    `faker:"email"`
	Password value.Password `faker:"len=8"`
	Name     value.Name     `faker:"len=3"`
}

// NewUser returns a new user instance.
func NewUser() *User {
	return &User{}
}

// String implements the fmt.Stringer interface.
func (User) String() string {
	return "user"
}

// Validate returns an error if the user is invalid.
func (u User) Validate() (*User, error) {
	var errs *errutil.Error

	if u.Email == "" {
		errs = errutil.Append(errs, errors.Required("Email"))
	}

	if u.Password == "" {
		errs = errutil.Append(errs, errors.Required("Password"))
	}

	if u.Name == "" {
		errs = errutil.Append(errs, errors.Required("Name"))
	}

	if errs.HasErrors() {
		return nil, errs
	}

	return &u, nil
}

// SetID sets the user ID.
func (u *User) SetID(id coreValue.ID) *User {
	u.ID = id
	return u
}

// SetEmail sets the user email.
func (u *User) SetEmail(email value.Email) *User {
	u.Email = email
	return u
}

// SetPassword sets the user password.
func (u *User) SetPassword(password value.Password) *User {
	u.Password = password
	return u
}

// SetName sets the user name.
func (u *User) SetName(name value.Name) *User {
	u.Name = name
	return u
}
