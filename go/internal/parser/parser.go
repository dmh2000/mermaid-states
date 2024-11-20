// Package parser deconstructs a mermaid state diagram graph into a structured format.
// It validates state names, transitions, and descriptions according to mermaid syntax rules.
package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	// placeholder is used when a transition has no description
	placeholder = "-"
	
	// maxLineLength is the maximum allowed length of an input line
	maxLineLength = 1000
	
	// maxInputLines is the maximum number of lines allowed in input
	maxInputLines = 10000
	
	// Regular expression patterns
	statePattern       = `^(?:[A-Za-z_][A-Za-z0-9_]*|\[\*\])$`
	transitionPattern  = `^([A-Za-z_][A-Za-z0-9_]*|\[\*\])\s*-->\s*([A-Za-z_][A-Za-z0-9_]*|\[\*\])(?:\s*\:(.+))?$`
	descriptionPattern = `^.+$`
)

// compile regular expressions once at package initialization
var (
	stateRegex       = regexp.MustCompile(statePattern)
	transitionRegex  = regexp.MustCompile(transitionPattern)
	descriptionRegex = regexp.MustCompile(descriptionPattern)
)

// Parser handles the parsing of mermaid state diagram syntax.
// It validates state names, transitions, and descriptions.
type Parser struct {
	stateRegex       *regexp.Regexp
	transitionRegex  *regexp.Regexp
	descriptionRegex *regexp.Regexp
}

// NewParser creates a new Parser instance with compiled regular expressions.
func NewParser() *Parser {
	return &Parser{
		stateRegex:       stateRegex,
		transitionRegex:  transitionRegex,
		descriptionRegex: descriptionRegex,
	}
}

// isValidState checks if a state name follows the required format.
func (p *Parser) isValidState(state string) bool {
	return p.stateRegex.MatchString(state)
}

// isValidTransition checks if a transition line follows the required format.
func (p *Parser) isValidTransition(line string) bool {
	return p.transitionRegex.MatchString(line)
}

// isValidDescription checks if a description is valid.
func (p *Parser) isValidDescription(desc string) bool {
	if desc == "" {
		return true
	}
	return p.descriptionRegex.MatchString(desc)
}

func (p *Parser) parseGraph(lines []string) ([]string, []string) {
	var validResults []string
	var invalidResults []string
	indent := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := p.transitionRegex.FindStringSubmatch(line)
		if matches != nil {
			fromState := matches[1]
			toState := matches[2]
			description := ""
			if len(matches) > 3 {
				description = matches[3]
			}

			if p.isValidState(fromState) && p.isValidState(toState) &&
				p.isValidDescription(description) {
				desc := description
				if description == "" {
					desc = placeholder
				}
				validResults = append(validResults,
					fmt.Sprintf("%s,%s,%s",
						fromState, toState, desc))
			}
		} else {
			invalidResults = append(invalidResults,
				fmt.Sprintf("%sInvalid input: %s",
					strings.Repeat("  ", indent), line))
		}
	}

	return validResults, invalidResults
}

func (p *Parser) parseInput(lines []string) ([]string, []string) {
	// Verify there is at least one transition
	hasValidTransition := false
	for _, line := range lines {
		if strings.TrimSpace(line) != "" && p.isValidTransition(strings.TrimSpace(line)) {
			hasValidTransition = true
			break
		}
	}

	if !hasValidTransition {
		return nil, []string{"Error: Graph must contain at least one transition"}
	}

	return p.parseGraph(lines)
}

// processInput reads lines from a scanner and parses them as a state diagram.
// It returns valid and invalid results, or an error if the input is invalid.
func processInput(scanner *bufio.Scanner) ([]string, []string, error) {
	var lines []string
	lineCount := 0

	// Set maximum line length
	scanner.Buffer(make([]byte, maxLineLength), maxLineLength)

	// Read all input lines
	for scanner.Scan() {
		lineCount++
		if lineCount > maxInputLines {
			return nil, nil, fmt.Errorf("input exceeds maximum of %d lines", maxInputLines)
		}
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading input: %w", err)
	}

	parser := NewParser()
	valid, invalid := parser.parseInput(lines)
	return valid, invalid, nil
}

// ProcessStateFile processes a state definition file and returns the valid results
// and an error if there are any invalid results. If verbose is true, invalid results
// are logged to stderr.
func ProcessStateFile(file *os.File, verbose bool) ([]string, error) {
	// Configure logging based on verbose flag
	if !verbose {
		log.SetOutput(io.Discard)
	} else {
		log.SetOutput(os.Stderr)
	}

	scanner := bufio.NewScanner(file)
	validResults, invalidResults, err := processInput(scanner)
	if err != nil {
		return nil, fmt.Errorf("processing input: %w", err)
	}

	// If there are invalid results, create an error with the details
	if len(invalidResults) > 0 {
		// Log invalid results if verbose
		for _, result := range invalidResults {
			if verbose {
				log.Println(result)
			}
		}
		return validResults, fmt.Errorf("found %d invalid state definitions", len(invalidResults))
	}

	return validResults, nil
}
