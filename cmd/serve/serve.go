package serve

import (
	"flag"
	"log"

	"github.com/smantic/cannonical/internal/server"
	"go.uber.org/zap"
)

// Run will start the http server.
func Run(args []string) {

	c := server.Config{}

	flags := flag.NewFlagSet("serve", flag.ExitOnError)
	flags.StringVar(&c.Addr, "address", "localhost", "address to run the server on")
	flags.StringVar(&c.Port, "port", "8080", "port to run the server on")
	flags.StringVar(&c.DebugPort, "debugport", "8081", "port for http server serving prom metrics and pprof to run on")

	err := flags.Parse(args)
	if err != nil {
		log.Printf("failed to parse flags %s\n", err.Error())
	}
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	c.Logger = logger

	s := server.NewServer(&c)

	log.Printf("grpc server running on %s:%s...\n", c.Addr, c.Port)
	err = s.Run()
	if err != nil {
		log.Printf("failed to run http server: %s\n", err.Error())
		return
	}
}
