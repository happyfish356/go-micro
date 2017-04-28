package rpc

import (
	"server"
)

func NewServer(opts ...server.Option) server.Server {
	return server.NewServer(opts...)
}
