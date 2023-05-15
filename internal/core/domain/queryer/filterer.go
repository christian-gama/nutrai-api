package queryer

type Filterer interface {
	// Field returns the field name.
	Field(idx int) string

	// Operator returns an operator string.
	Operator(idx int) string

	// Value returns a value string.
	Value(idx int) string

	// Slice returns a slice of strings.
	Slice() []string

	// Add adds a filter to the filterer.
	Add(field string, operator string, value any) Filterer
}
