package stategen

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

type TestCase struct {
	Name        string   `json:"name"`
	Input       []string `json:"input"`
	WantValid   []string `json:"wantValid"`
	WantInvalid []string `json:"wantInvalid"`
}

type TestSuite struct {
	Tests []TestCase `json:"tests"`
}

func loadTestCases(t *testing.T) []TestCase {
	// Get the absolute path to the test file
	absPath, err := filepath.Abs("../../test/tests.json")
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	// Read the test file
	data, err := os.ReadFile(absPath)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Parse the JSON
	var suite TestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("Failed to parse test file: %v", err)
	}

	return suite.Tests
}

func TestParser_ParseGraph(t *testing.T) {
	parser := NewParser()
	tests := loadTestCases(t)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			gotValid, gotInvalid := parser.parseGraph(tt.Input)
			if !reflect.DeepEqual(gotValid, tt.WantValid) {
				t.Errorf("Parser.parseGraph() valid = %v, want %v", gotValid, tt.WantValid)
			}
			if !reflect.DeepEqual(gotInvalid, tt.WantInvalid) {
				t.Errorf("Parser.parseGraph() invalid = %v, want %v", gotInvalid, tt.WantInvalid)
			}
		})
	}
}

func TestWithInputFile(t *testing.T) {
	// Read input file
	inputFile, err := os.Open("../../test/input.txt")
	if err != nil {
		t.Fatalf("Failed to open input file: %v", err)
	}
	defer inputFile.Close()

	// Read output file
	outputFile, err := os.Open("../../test/output.txt")
	if err != nil {
		t.Fatalf("Failed to open output file: %v", err)
	}
	defer outputFile.Close()

	// Read expected output lines
	var expectedOutput []string
	scanner := bufio.NewScanner(outputFile)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			expectedOutput = append(expectedOutput, line)
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	// Process input file
	var inputLines []string
	scanner = bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			inputLines = append(inputLines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("Error reading input file: %v", err)
	}

	// Parse input and compare with expected output
	parser := NewParser()
	gotValid, gotInvalid := parser.parseGraph(inputLines)

	if len(gotInvalid) > 0 {
		t.Errorf("Got unexpected invalid lines: %v", gotInvalid)
	}

	if !reflect.DeepEqual(gotValid, expectedOutput) {
		t.Errorf("Parser output doesn't match expected output\nGot: %v\nWant: %v", gotValid, expectedOutput)
	}
}

func TestInvalidInputs(t *testing.T) {
	parser := NewParser()

	invalidInputs := []struct {
		name        string
		input       []string
		wantInvalid []string
	}{
		{
			name:        "invalid state characters",
			input:       []string{"State#1 --> State2"},
			wantInvalid: []string{"Invalid input: State#1 --> State2"},
		},
		{
			name:        "missing arrow",
			input:       []string{"State1 State2"},
			wantInvalid: []string{"Invalid input: State1 State2"},
		},
		{
			name:        "wrong arrow direction",
			input:       []string{"State1 <-- State2"},
			wantInvalid: []string{"Invalid input: State1 <-- State2"},
		},
	}

	for _, tt := range invalidInputs {
		t.Run(tt.name, func(t *testing.T) {
			_, gotInvalid := parser.parseGraph(tt.input)
			if !reflect.DeepEqual(gotInvalid, tt.wantInvalid) {
				t.Errorf("Parser.parseGraph() invalid = %v, want %v", gotInvalid, tt.wantInvalid)
			}
		})
	}
}
