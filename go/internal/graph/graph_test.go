package graph

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const testdir = "../../../test/"

func TestParseEdge(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantEdge *Edge
	}{
		{
			name:  "valid edge",
			input: "a,b,c",
			wantEdge: &Edge{
				From:        "a",
				To:          "b",
				Description: "c",
			},
		},
		{
			name:     "empty string",
			input:    "",
			wantEdge: nil,
		},
		{
			name:     "wrong number of fields",
			input:    "a,b",
			wantEdge: nil,
		},
		{
			name:     "empty from field",
			input:    ",b,c",
			wantEdge: nil,
		},
		{
			name:     "empty to field",
			input:    "a,,c",
			wantEdge: nil,
		},
		{
			name: "whitespace handling",
			input: " a , b , c ",
			wantEdge: &Edge{
				From:        "a",
				To:          "b",
				Description: "c",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseEdge(tt.input)
			if (got == nil) != (tt.wantEdge == nil) {
				t.Errorf("ParseEdge() = %v, want %v", got, tt.wantEdge)
				return
			}
			if got != nil {
				if got.From != tt.wantEdge.From || got.To != tt.wantEdge.To || got.Description != tt.wantEdge.Description {
					t.Errorf("ParseEdge() = %v, want %v", got, tt.wantEdge)
				}
			}
		})
	}
}

func TestGraph(t *testing.T) {
	g := NewGraph()
	
	// Test empty graph
	if len(g.Nodes) != 0 {
		t.Error("new graph should be empty")
	}
	
	// Test adding nodes
	g.AddNode("a")
	if len(g.Nodes["a"]) != 0 {
		t.Error("new node should have no edges")
	}
	
	// Test adding same node twice
	g.AddNode("a")
	if len(g.Nodes) != 1 {
		t.Error("adding same node twice should not create duplicate")
	}
	
	// Test adding edge
	edge := &Edge{From: "a", To: "b", Description: "test"}
	g.AddEdge(edge)
	
	if len(g.Nodes["a"]) != 1 {
		t.Error("edge not added to from node")
	}
	
	if _, exists := g.Nodes["b"]; !exists {
		t.Error("to node not created")
	}
	
	// Test nil edge
	g.AddEdge(nil)
	if len(g.Nodes["a"]) != 1 {
		t.Error("nil edge should not be added")
	}
}

func TestLoad1(t *testing.T) {
	g := NewGraph()
	g.Load([]string{"a,b,c", "d,e,f"})
	if g.Nodes["a"][0].From != "a" || g.Nodes["a"][0].To != "b" || g.Nodes["a"][0].Description != "c" {
		t.Errorf("expected edge {a,b,c}, got %v", g.Nodes["a"][0])
	}
	if g.Nodes["d"][0].From != "d" || g.Nodes["d"][0].To != "e" || g.Nodes["d"][0].Description != "f" {
		t.Errorf("expected edge {d,e,f}, got %v", g.Nodes["d"][0])
	}
}

func TestLoad2(t *testing.T) {

	// read the test/graph.txt file and convert it to a slice of strings
	lines, err := os.ReadFile(fmt.Sprintf("%s%s", testdir, "graph.txt"))
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	g := NewGraph()
	g.Load(strings.Split(string(lines), "\n"))

	t.Log(g)
}
