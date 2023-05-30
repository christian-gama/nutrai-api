package value

// To represents the mail receiver.
type To struct {
	Email string `json:"email" faker:"email"`
	Name  string `json:"name" faker:"name"`
}

// NewTo creates a new To.
func NewTo() *To {
	return &To{}
}

// SetEmail sets the Email field.
func (t *To) SetEmail(email string) *To {
	t.Email = email
	return t
}

// SetName sets the Name field.
func (t *To) SetName(name string) *To {
	t.Name = name
	return t
}
