package main

import (
	"flag"
	"fmt"
	"os"

	_ "net/http/pprof"

	"github.com/smantic/cannonical/cmd"
	"github.com/smantic/cannonical/cmd/serve"
)

func main() {

	flag.Usage = func() {
		fmt.Printf(cmd.HelpStr)
	}

	if len(os.Args) == 1 {
		fmt.Printf(cmd.HelpStr)
		return
	}

	switch os.Args[1] {
	// print help message
	case "help":
		fmt.Printf(cmd.HelpStr)
		return
	// serve an http server
	case "serve":
		serve.Run(os.Args[2:])
	}

	flag.Parse()
}
