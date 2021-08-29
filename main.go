package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/smantic/cannonical/cmd"
	"github.com/smantic/cannonical/cmd/serve"
)

func init() {
	// set flags

	flag.Usage = func() {
		fmt.Printf(cmd.HelpStr)
	}
}

func main() {

	if len(os.Args) == 1 {
		fmt.Printf(cmd.HelpStr)
		return
	}

	switch os.Args[1] {
	case "help":
		fmt.Printf(cmd.HelpStr)
		return
	case "list":
		fmt.Printf(cmd.CommandStr)
		return
	// serve an http server
	case "serve":
		serve.Run(os.Args[2:])
	}

	flag.Parse()
}
