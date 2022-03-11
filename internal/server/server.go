package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Config is the config for the server.
type Config struct {
	Addr      string
	Port      string
	DebugPort string

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

	lis, err := net.Listen("tcp", s.Addr+":"+s.Port)
	if err != nil {
		return err
	}

	grpc_zap.ReplaceGrpcLoggerV2(s.Logger)
	g := grpc.NewServer(
		middleware.WithUnaryServerChain(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(s.Logger),
		),
		middleware.WithStreamServerChain(
			grpc_prometheus.StreamServerInterceptor,
			grpc_ctxtags.StreamServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
			),
			grpc_zap.StreamServerInterceptor(s.Logger),
		),
	)

	grpc_prometheus.Register(g)
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("http server running on :%s...\n", s.DebugPort)
		err := http.ListenAndServe(":"+s.DebugPort, http.DefaultServeMux)
		if err != nil {
			fmt.Printf("failed to mount http server; %s", err.Error())
		}
	}()

	return g.Serve(lis)
}
