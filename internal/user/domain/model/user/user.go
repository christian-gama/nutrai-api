package user

import (
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// User is the user model.
type User struct {
	ID       sharedvalue.ID `faker:"uint"`
	Email    value.Email    `faker:"email"`
	Password value.Password `faker:"len=8"`
	Name     value.Name     `faker:"len=3"`
}

// Validate returns an error if the user is invalid.
func (u *User) Validate() error {
	var errs *errutil.Error

	if err := u.Email.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := u.Password.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := u.Name.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}

type builder struct {
	user *User
}

// NewBuilder returns a new builder for the user model.
func NewBuilder() *builder {
	return &builder{
		user: &User{},
	}
}

// SetID sets the user ID.
func (b *builder) SetID(id sharedvalue.ID) *builder {
	b.user.ID = id
	return b
}

// SetEmail sets the user email.
func (b *builder) SetEmail(email value.Email) *builder {
	b.user.Email = email
	return b
}

// SetPassword sets the user password.
func (b *builder) SetPassword(password value.Password) *builder {
	b.user.Password = password
	return b
}

// SetName sets the user name.
func (b *builder) SetName(name value.Name) *builder {
	b.user.Name = name
	return b
}

// Build returns the user model.
func (b *builder) Build() (*User, error) {
	if err := b.user.Validate(); err != nil {
		return nil, err
	}

	return b.user, nil
}
