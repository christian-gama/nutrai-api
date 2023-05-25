package validator

type Validator interface {
	// Validate validates the given struct.
	Validate(anStruct any) error
}
