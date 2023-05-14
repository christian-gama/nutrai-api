// Exposes the restricted set of HTTP methods instead of the full set of HTTP methods
// to keep the API standardized and simple.
package http

const (
	// MethodGet is the HTTP GET method.
	MethodGet = "GET"

	// MethodPost is the HTTP POST method.
	MethodPost = "POST"

	// MethodPut is the HTTP PUT method.
	MethodPut = "PUT"

	// MethodDelete is the HTTP DELETE method.
	MethodDelete = "DELETE"
)
