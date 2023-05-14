// Exposes the restricted set of HTTP methods instead of the full set of HTTP methods
// to keep the API standardized and simple.
package http

type Method string

const (
	// MethodGet is the HTTP GET method.
	MethodGet Method = "GET"

	// MethodPost is the HTTP POST method.
	MethodPost Method = "POST"

	// MethodPut is the HTTP PUT method.
	MethodPut Method = "PUT"

	// MethodDelete is the HTTP DELETE method.
	MethodDelete Method = "DELETE"
)
