#!/usr/bin/env python3
"""
State Machine Graph Parser

This module provides functionality to parse and validate state machine transition
definitions in a simple text format. It supports states with alphanumeric names
and transitions between states with optional descriptions.

Format:
    - States: Alphanumeric strings starting with a letter/underscore, or [*]
    - Transitions: state1 --> state2 : description
    - Description: Optional text following the transition

Example:
    START --> PROCESSING : begin
    PROCESSING --> END : complete
"""

import sys
import re
import logging
import argparse

# Default placeholder for empty transition descriptions
PLACEHOLDER = "-"

class Parser:
    """
    Parser for state machine transition definitions.
    
    Validates and parses state transitions according to the following grammar:
    - STATE: [A-Za-z_][A-Za-z0-9_]* | [*]
    - TRANSITION: STATE --> STATE [:DESCRIPTION]
    - DESCRIPTION: Any non-empty string
    """
    
    def __init__(self):
        # Define regex patterns based on BNF grammar
        # Matches valid state names: letters, numbers, underscore (must start with letter/underscore) or [*]
        self.state_regex = re.compile(r'^(?:[A-Za-z_][A-Za-z0-9_]*|\[\*\])$')
        # Matches transitions with optional description: state1 --> state2 : description
        self.transition_regex = re.compile(r'^([A-Za-z_][A-Za-z0-9_]*|\[\*\])\s*-->\s*([A-Za-z_][A-Za-z0-9_]*|\[\*\])(?:\s*\:(.+))?$')
        # Matches any non-empty description
        self.description_regex = re.compile(r'^.+$')

    def is_valid_state(self, state):
        """
        Validate if a string is a valid state name.
        
        Args:
            state (str): The state name to validate
            
        Returns:
            bool: True if the state name is valid, False otherwise
        """
        return bool(self.state_regex.match(state))
    
    def is_valid_transition(self, line):
        """
        Validate if a string represents a valid transition.
        
        Args:
            line (str): The transition definition to validate
            
        Returns:
            bool: True if the transition is valid, False otherwise
        """
        return bool(self.transition_regex.match(line))

    def is_valid_description(self, desc):
        """
        Validate if a string is a valid transition description.
        
        Args:
            desc (str): The description to validate
            
        Returns:
            bool: True if the description is valid, False otherwise
        """
        logging.debug(desc)
        if not desc:
            return True
        return bool(self.description_regex.match(desc))
    
    def parse_graph(self, lines):
        """
        Parse a graph consisting of one or more transitions.
        
        Args:
            lines (list): List of strings containing transition definitions
            
        Returns:
            tuple: (valid_results, invalid_results) where:
                  - valid_results is a list of validated transitions in CSV format
                  - invalid_results is a list of error messages for invalid inputs
        """
        valid_results = []
        invalid_results = []

        for line in lines:
            line = line.strip()
            if not line:
                continue
                
            matches = self.transition_regex.match(line)
            if matches:
                from_state = matches.group(1)
                to_state = matches.group(2)
                description = ""
                if len(matches.groups()) > 2 and matches.group(3):
                    description = matches.group(3).strip()

                if self.is_valid_state(from_state) and self.is_valid_state(to_state) and self.is_valid_description(description):
                    desc = description if description else PLACEHOLDER
                    valid_results.append(f"{from_state},{to_state},{desc}")
                else:
                    invalid_results.append(f"Invalid input: {line}")
            else:
                invalid_results.append(f"Invalid input: {line}")

        return valid_results, invalid_results

    def parse_input(self, lines):
        """
        Parse and validate all input lines as a state machine graph.
        
        Args:
            lines (list): List of strings containing transition definitions
            
        Returns:
            tuple: (valid_results, invalid_results) where:
                  - valid_results is a list of validated transitions in CSV format
                  - invalid_results is a list of error messages for invalid inputs
        """
        # Verify there is at least one transition
        has_valid_transition = False
        for line in lines:
            if line.strip() and self.is_valid_transition(line.strip()):
                has_valid_transition = True
                break

        if not has_valid_transition:
            return [], ["Error: Graph must contain at least one transition"]
            
        return self.parse_graph(lines)

def main():
    """
    Main entry point for the state machine graph parser.
    
    Reads input from a file or stdin, parses the state transitions,
    and outputs valid transitions in CSV format to stdout and errors to stderr.
    """
    arg_parser = argparse.ArgumentParser(description='Parse state transition definitions')
    arg_parser.add_argument('file', nargs='?', type=argparse.FileType('r'), default=sys.stdin,
                        help='Input file (if not specified, reads from stdin)')
    arg_parser.add_argument('-v', '--verbose', action='store_true',
                        help='Enable verbose logging output')

    args = arg_parser.parse_args()
    
    # Configure logging based on verbose flag
    if args.verbose:
        logging.basicConfig(format='%(message)s', level=logging.ERROR)
    else:
        logging.basicConfig(format='%(message)s', level=logging.CRITICAL)  # Suppress most logging
    
    try:
        lines = [line.rstrip() for line in args.file]
    except Exception as e:
        logging.error(f"Error reading input: {e}")
        sys.exit(1)
    finally:
        if args.file is not sys.stdin:
            args.file.close()

    if not lines:
        logging.error("No input provided")
        sys.exit(1)

    parser = Parser()
    valid_results, invalid_results = parser.parse_input(lines)
    
    # Print valid results to stdout
    for result in valid_results:
        print(result)
        
    # Print invalid results to stderr
    for result in invalid_results:
        if args.verbose:
            logging.error(result)

if __name__ == "__main__":
    main()
