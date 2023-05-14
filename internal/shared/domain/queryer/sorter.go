package queryer

type Sorter interface {
	// Field returns the field name.
	Field(idx int) string

	// IsDesc returns a boolean indicating whether the sort is descending.
	IsDesc(idx int) bool

	// Slice returns a slice of strings.
	Slice() []string

	// Add adds a sort to the sorter.
	Add(field string, isDesc bool) Sorter
}
