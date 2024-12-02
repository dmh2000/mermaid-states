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
	var input = os.Stdin
	var err error

	// If there's a non-flag argument, treat it as input file
	if flag.NArg() > 0 {
		filename := flag.Arg(0)
		input, err = os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer input.Close()
	}

	// Process input using the parser
	validResults, err := parser.ProcessStateFile(input, *verbose)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing input: %v\n", err)
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
