package main

import (
	"flag"
	"fmt"
	"os"

	stategen "sqirvy.xyz/state-gen/state-gen"
)

func main() {
	// Define command line flags
	verbose := flag.Bool("v", false, "Enable verbose logging output")
	flag.Parse()

	exitCode := 0

	// Process stdin using the new function
	validResults, err := stategen.ProcessStateFile(os.Stdin, *verbose)
	if err != nil {
		exitCode = 1
	}

	// Print valid results to stdout
	for _, result := range validResults {
		fmt.Println(result)
	}

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}
