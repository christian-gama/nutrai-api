package response

// Body is the response body of the API. It contains the status of the
// request, the data and the stack trace (if any). The stack trace is only
// returned when the application is in debug mode (CONFIG_DEBUG=true).
// Data will be omitted if it's nil.
type Body struct {
	Status bool `json:"status"`
	Data   any  `json:"data,omitempty"`
	Stack  any  `json:"stack,omitempty"`
}
