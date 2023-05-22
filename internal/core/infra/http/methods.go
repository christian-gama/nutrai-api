// Exposes the restricted set of HTTP methods instead of the full set of HTTP methods
// to keep the API standardized and simple.
package http

// Method is the HTTP method.
type Method string

const (
	// MethodGet is the HTTP GET method.
	MethodGet Method = "GET"

	// MethodPost is the HTTP POST method.
	MethodPost Method = "POST"

	// MethodPut is the HTTP PUT method.
	MethodPut Method = "PUT"

	// MethodPatch is the HTTP PATCH method.
	MethodPatch Method = "PATCH"

	// MethodDelete is the HTTP DELETE method.
	MethodDelete Method = "DELETE"
)

// String returns the string representation of the HTTP method.
func (m Method) String() string {
	return string(m)
}
