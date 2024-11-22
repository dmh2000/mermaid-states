// Package graph creates the graph of states and transitions
// this package inputs the .csv file produced by the parser, and
// builds a directed graph of states and transitions. this graph
// will be used to create the states for the state machine
package graph

import (
	"fmt"
	"regexp"
	"strings"
)

var edgePattern = regexp.MustCompile(`^([^,]+),([^,]+),(.*)$`)

// Edge represents a directed edge in the graph with a description
type Edge struct {
	// From is the source node
	From string
	// To is the destination node
	To string
	// Description is the label or description of the edge
	Description string
}

// ParseEdge parses a comma-separated string into an Edge
// Format: "from,to,description"
func ParseEdge(t string) (*Edge, error) {
	if t == "" {
		return nil, fmt.Errorf("syntax: empty edge")
	}

	// use regex to capture the three parts
	matches := edgePattern.FindStringSubmatch(t)
	if matches == nil || len(matches) != 4 {
		return nil, fmt.Errorf("syntax: invalid edge format %v", t)
	}

	// validate fields (matches[0] is the full match)
	from := strings.TrimSpace(matches[1])
	to := strings.TrimSpace(matches[2])
	desc := strings.TrimSpace(matches[3])

	if from == "" || to == "" {
		return nil, fmt.Errorf("syntax: invalid fromt/to format %v", t)
	}

	return &Edge{
			From:        from,
			To:          to,
			Description: desc,
		},
		nil
}

// Graph represents a directed graph using an adjacency list
type Graph struct {
	// Nodes maps node names to their outgoing edges
	Nodes map[string][]Edge
}

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string][]Edge),
	}
}

func (g *Graph) AddNode(node string) {
	_, ok := g.Nodes[node]
	// is it already in the graph
	if ok {
		return
	}
	// if not, add it with an empty edge list
	g.Nodes[node] = make([]Edge, 0)
}

func (g *Graph) AddEdge(edge *Edge) {
	if edge == nil {
		return
	}

	// ensure the from node exists
	g.AddNode(edge.From)

	// add the edge
	g.Nodes[edge.From] = append(g.Nodes[edge.From], *edge)

	// ensure the to node exists
	g.AddNode(edge.To)
}

func (g *Graph) Load(s []string) error {
	for _, t := range s {
		edge, err := ParseEdge(t)
		if err != nil {
			return err
		}
		if edge != nil {
			g.AddEdge(edge)
		}
	}
	return nil
}

func (g *Graph) String() string {
	var sb strings.Builder
	for node, edges := range g.Nodes {
		sb.WriteString("-------\nnode: ")
		sb.WriteString(node)
		sb.WriteString("\n")
		for _, edge := range edges {
			sb.WriteString("    ")
			sb.WriteString(edge.From)
			sb.WriteString(" -> ")
			sb.WriteString(edge.To)
			sb.WriteString(" : ")
			sb.WriteString(edge.Description)
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
