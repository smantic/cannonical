package server

import (
	"net"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Config is the config for the server.
type Config struct {
	Addr string
	Port string

	// Debug tells us if we should mount pprof
	Debug bool

	Logger *zap.Logger
}

// Server is our http server.
type Server struct {
	Config
}

// NewServer will create a new server.
func NewServer(c *Config) Server {

	return Server{
		Config: *c,
	}
}

// Run will run our http server.
func (s *Server) Run() error {

	lis, err := net.Listen("tcp", s.Port)
	if err != nil {
		return err
	}

	grpc_zap.ReplaceGrpcLoggerV2(s.Logger)
	g := grpc.NewServer(
		middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(s.Logger),
		),
		middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(s.Logger),
		),
	)

	return g.Serve(lis)
}
