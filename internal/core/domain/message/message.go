package message

// MessageHandler is the message handler function.
type MessageHandler func(body []byte) error

// Consumer is the consumer interface.
type Consumer interface {
	Handle(handler MessageHandler)
}
