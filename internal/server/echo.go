package server

import (
	"context"

	"github.com/smantic/cannonical/proto"
)

type EchoServer struct {
	proto.UnimplementedEchoServer
}

func (e *EchoServer) Echo(ctx context.Context, req *proto.StringMessage) (*proto.StringMessage, error) {

	msg := proto.StringMessage{
		Value: req.Value,
	}

	return &msg, nil
}
