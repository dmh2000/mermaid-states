package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const testdir = "../../../test/"

// TestCase represents a single test case for the parser
type TestCase struct {
	Name        string   `json:"name"`
	Input       []string `json:"input"`
	WantValid   []string `json:"wantValid"`
	WantInvalid []string `json:"wantInvalid"`
}

// TestSuite represents a collection of test cases
type TestSuite struct {
	Tests []TestCase `json:"tests"`
}

// loadTestCases loads test cases from the JSON file
func loadTestCases(t *testing.T) []TestCase {
	t.Helper()

	// Get the absolute path to the test file
	relPath := filepath.Clean(fmt.Sprintf("%s%s", testdir, "tests.json"))

	// Verify the path is under the test directory
	if !strings.HasPrefix(relPath, filepath.Clean(testdir)) {
		t.Fatalf("Test file path escapes test directory")
	}

	// Read the test file
	data, err := os.ReadFile(relPath)
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
	// Clean and verify input file path
	inputPath := filepath.Clean(fmt.Sprintf("%s%s", testdir, "input.txt"))
	if !strings.HasPrefix(inputPath, filepath.Clean(testdir)) {
		t.Fatal("Input file path escapes test directory")
	}

	// Clean and verify output file path
	outputPath := filepath.Clean(fmt.Sprintf("%s%s", testdir, "output.txt"))
	if !strings.HasPrefix(outputPath, filepath.Clean(testdir)) {
		t.Fatal("Output file path escapes test directory")
	}

	// Read input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		t.Fatalf("Failed to open input file: %v", err)
	}
	defer inputFile.Close()

	// Read output file
	outputFile, err := os.Open(outputPath)
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
	validResults, err := ProcessStateFile(inputFile, false)
	if err != nil {
		t.Fatalf("ProcessStateFile failed: %v", err)
	}

	if !reflect.DeepEqual(validResults, expectedOutput) {
		t.Errorf("Parser output doesn't match expected output\nGot: %v\nWant: %v", validResults, expectedOutput)
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

func TestSecurityLimits(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "line too long",
			input:   strings.Repeat("a", maxLineLength+1),
			wantErr: true,
		},
		{
			name:    "too many lines",
			input:   strings.Repeat("State1 --> State2\n", maxInputLines+1),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			scanner := bufio.NewScanner(r)
			_, _, err := processInput(scanner)
			if (err != nil) != tt.wantErr {
				t.Errorf("processInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
