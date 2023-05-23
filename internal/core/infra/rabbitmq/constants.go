package rabbitmq

import "github.com/rabbitmq/amqp091-go"

type Exchange = string

const (
	ExchangeDirect  Exchange = amqp091.ExchangeDirect
	ExchangeFanout  Exchange = amqp091.ExchangeFanout
	ExchangeTopic   Exchange = amqp091.ExchangeTopic
	ExchangeHeaders Exchange = amqp091.ExchangeHeaders
)

type ContentType = string

const (
	ContentTypeTextPlain ContentType = "text/plain"
	ContentTypeJSON      ContentType = "application/json"
	ContentTypeXML       ContentType = "application/xml"
	ContentTypeProtobuf  ContentType = "application/protobuf"
)
