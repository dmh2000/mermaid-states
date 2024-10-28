package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Parser struct {
	stateRegex       *regexp.Regexp
	transitionRegex  *regexp.Regexp
	descriptionRegex *regexp.Regexp
}

func NewParser() *Parser {
	return &Parser{
		stateRegex:       regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`),
		transitionRegex:  regexp.MustCompile(`^([A-Za-z_][A-Za-z0-9_]*)\s*-->\s*([A-Za-z_][A-Za-z0-9_]*)(?:\s*:\s*(.+))?$`),
		descriptionRegex: regexp.MustCompile(`^[ a-zA-Z0-9_\-:.,?!@=~]+$`),
	}
}

func (p *Parser) isValidState(state string) bool {
	return p.stateRegex.MatchString(state)
}

func (p *Parser) isValidTransition(line string) bool {
	return p.transitionRegex.MatchString(line)
}

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
				descText := ""
				if description != "" {
					descText = ": " + description
				}
				validResults = append(validResults,
					fmt.Sprintf("%sValid transition from %s to %s%s",
						strings.Repeat("  ", indent), fromState, toState, descText))
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	// Read all input lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	parser := NewParser()
	validResults, invalidResults := parser.parseInput(lines)

	// Print valid results to stdout
	for _, result := range validResults {
		fmt.Println(result)
	}

	// Print invalid results to stderr
	for _, result := range invalidResults {
		fmt.Fprintln(os.Stderr, result)
	}
}
