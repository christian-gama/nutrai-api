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

// New returns a new User instance.
func New(input *UserInput) (*User, error) {
	user := User(*input)

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return &user, nil
}

// Validate returns an error if the user is invalid.
func (u *User) Validate() error {
	var errs *errutil.Error

	if err := u.ID.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

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
