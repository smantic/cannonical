package server

import (
	"context"

	"github.com/smantic/cannonical/proto"
)

func (s *Server) Check(context.Context, *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	return &proto.HealthCheckResponse{
		Status: proto.HealthCheckResponse_SERVING,
	}, nil
}

func (s *Server) Watch(req *proto.HealthCheckRequest, server proto.Health_WatchServer) error {
	return server.Send(&proto.HealthCheckResponse{
		Status: proto.HealthCheckResponse_SERVING,
	})
}
