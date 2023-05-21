package message

// Publisher is the publisher interface.
type Publisher interface {
	Handle(msg []byte)
}
