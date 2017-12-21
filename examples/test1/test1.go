// [flag package](https://golang.org/pkg/flag/) is the Go equivalent
// of Python [argparse](https://docs.python.org/2/howto/argparse.html).
// While not as powerful, it does what we expect it to do. It
// simplifies adding and parsing command line parameters, leaving us
// to concentrate on the tools. Most of our tools will need them to be
// actually useful (hardcoding URLs and IPs get old too fast).

package main

import (
	"flag"
	"fmt"
)

func main() {

	// Declare flags
	// Remember, flag methods return pointers
	ipPtr := flag.String("ip", "127.0.0.1", "target IP")

	var port int
	flag.IntVar(&port, "port", 8080, "Port")

	verbosePtr := flag.Bool("verbose", true, "verbosity")

	// Parse flags
	flag.Parse()

	// Hack IP:port
	fmt.Printf("Hacking %s:%d!\n", *ipPtr, port)

	// Display progress if verbose flag is set
	if *verbosePtr {
		fmt.Printf("Pew pew!\n")
	}
}
