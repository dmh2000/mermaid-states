// Package graph creates the graph of states and transitions
// this package inputs the .csv file produced by the parser, and
// builds a directed graph of states and transitions. this graph
// will be used to create the states for the state machine
package graph

import (
	"strings"
)

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
func ParseEdge(t string) *Edge {
	if t == "" {
		return nil
	}

	// split the input string into from, to, description strings separated by commas
	fields := strings.Split(t, ",")
	if len(fields) != 3 {
		return nil
	}

	// validate fields
	from := strings.TrimSpace(fields[0])
	to := strings.TrimSpace(fields[1])
	desc := strings.TrimSpace(fields[2])

	if from == "" || to == "" {
		return nil
	}

	return &Edge{
		From:        from,
		To:          to,
		Description: desc,
	}
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

func (g *Graph) Load(s []string) {
	for _, t := range s {
		edge := ParseEdge(t)
		if edge != nil {
			g.AddEdge(edge)
		}
	}
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
