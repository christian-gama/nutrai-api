package event

// Action is the action of an event.
type Action string

const (
	Save   Action = "save"
	Update Action = "update"
	Delete Action = "delete"
	Error  Action = "error"
)
