package http

import (
	"transport"
)

func NewTransport(opts ...transport.Option) transport.Transport {
	return transport.NewTransport(opts...)
}
