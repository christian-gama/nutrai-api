package queryer

type Paginator interface {
	// GetLimit returns a limit which indicates the maximum number of results to be returned.
	GetLimit() int

	// GetPage returns a page number.
	GetPage() int

	// GetOffset returns a number of results to skip.
	GetOffset() int

	// SetLimit sets a limit which indicates the maximum number of results to be returned.
	SetLimit(limit int) Paginator

	// SetPage sets a page number.
	SetPage(page int) Paginator
}

// PaginationOutput is a struct which contains the total number of results and a slice of results.
type PaginationOutput[Result any] struct {
	Total   int
	Results []Result
}
