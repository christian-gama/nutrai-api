package message

// MessageHandler is the message handler function.

// Consumer is the consumer interface.
type Consumer[Data any] interface {
	Handle(handler func(data Data) error)
}
