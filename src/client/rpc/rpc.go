package rpc

import (
	"client"
)

func NewClient(opts ...client.Option) client.Client {
	return client.NewClient(opts...)
}
