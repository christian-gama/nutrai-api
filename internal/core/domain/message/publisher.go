package message

// Publisher is the publisher interface.
type Publisher[Data any] interface {
	Handle(data Data)
}
