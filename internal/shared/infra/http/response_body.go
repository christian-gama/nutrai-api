package http

type ResponseBody struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
	Stack  any  `json:"stack,omitempty"`
}
