package value

// Subject represents the subject of an email message.
type Subject string

// String returns the string representation of the subject.
func (s Subject) String() string {
	return string(s)
}
