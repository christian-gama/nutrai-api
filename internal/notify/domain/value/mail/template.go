package value

// Template represents the template of an email message.
type Template string

// String returns the string representation of the template.
func (t Template) String() string {
	return string(t)
}
