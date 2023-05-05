package value

// ID represents the ID of a domain entity.
type ID uint

// Uint returns the ID as an unsigned integer.
func (i ID) Uint() uint {
	return uint(i)
}
