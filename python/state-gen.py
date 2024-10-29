from sys import stdin, stderr
import re
import logging

PLACEHOLDER = "-"

class Parser:
    def __init__(self):
        # Define regex patterns based on BNF grammar
        self.state_regex = re.compile(r'^(?:[A-Za-z_][A-Za-z0-9_]+|\[\*\])$')
        self.transition_regex = re.compile(r'^([A-Za-z_][A-Za-z0-9_]*|\[\*\])\s*-->\s*([A-Za-z_][A-Za-z0-9_]*|\[\*\])(?:\s*\:(.+))?$')
        self.description_regex = re.compile(r'^.+$')

    def is_valid_state(self, state):
        """Check if a string matches the STATE rule"""
        return bool(self.state_regex.match(state))
    
    def is_valid_transition(self, line):
        """Check if a string matches the TRANSITION rule"""
        return bool(self.transition_regex.match(line))

    def is_valid_description(self, desc):
        """Check if a string matches the DESCRIPTION rule"""
        logging.debug(desc)
        if not desc:
            return True
        return bool(self.description_regex.match(desc))
    
    def parse_graph(self, lines):
        """Parse a graph consisting of one or more transitions"""
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

        return valid_results, invalid_results

    def parse_input(self, lines):
        """Parse all input as a graph of transitions"""
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
    lines = []
    for line in stdin:
        lines.append(line.rstrip())

    parser = Parser()
    valid_results, invalid_results = parser.parse_input(lines)
    
    # Print valid results to stdout
    for result in valid_results:
        print(result)
        
    # Print invalid results to stderr
    for result in invalid_results:
        logging.error(result)

if __name__ == "__main__":
    logging.basicConfig(format='%(message)s', level=logging.ERROR)
    main()
