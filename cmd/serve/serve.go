package serve

import (
	"flag"
	"log"

	"github.com/smantic/cannonical/internal/server"
)

func flags(c *server.Config) {

}

// Run will start the http server
func Run(args []string) {

	c := server.Config{}

	flags := flag.NewFlagSet("serve", flag.ExitOnError)
	flags.StringVar(&c.Addr, "address", "localhost", "address to run the server on")
	flags.StringVar(&c.Port, "port", "8080", "port to run the server on")
	flags.Parse(args)

	log.Printf("http server running on %s:%s...\n", c.Addr, c.Port)
	err := server.Serve(c)
	if err != nil {
		log.Printf("failed to run http server: %s\n", err.Error())
		return
	}
}
