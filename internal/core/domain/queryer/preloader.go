package queryer

type Preloader interface {
	// Slice returns a slice of strings.
	Slice() []string

	// Add adds a preload to the Preloader.
	Add(field string) Preloader
}
