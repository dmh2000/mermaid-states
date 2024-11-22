package main

import (
	"flag"
	"fmt"
	"os"

	graph "sqirvy.xyz/state-gen/internal/graph"
	parser "sqirvy.xyz/state-gen/internal/parser"
)

func main() {
	// Define command line flags
	verbose := flag.Bool("v", false, "Enable verbose logging output")
	flag.Parse()

	exitCode := 0

	// Process stdin using the new function
	validResults, err := parser.ProcessStateFile(os.Stdin, *verbose)
	if err != nil {
		exitCode = 1
	}

	// Print valid results to stdout
	g := graph.NewGraph()
	err = g.Load(validResults)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Load Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(g)

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}
