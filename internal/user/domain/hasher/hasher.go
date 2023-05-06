package hasher

type Hasher interface {
	// Hash returns the hashed value.
	Hash(value string) (string, error)

	// Compare returns an error if the value is not equal to the hashed value.
	Compare(value, hash string) error
}
