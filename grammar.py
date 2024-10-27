import sys
import re

class Parser:
    def __init__(self):
        # Define regex patterns based on BNF grammar
        self.state_pattern = r'[A-Za-z*][A-Za-z0-9_]*'
        self.transition_pattern = rf'{self.state_pattern}\s*-->\s*{self.state_pattern}'
        
    def is_valid_state(self, state):
        """Check if a string matches the STATE rule"""
        return bool(re.match(f'^{self.state_pattern}$', state))
    
    def is_valid_transition(self, line):
        """Check if a string matches the TRANSITION rule"""
        return bool(re.match(f'^{self.transition_pattern}$', line))
    
    def parse_line(self, line):
        """Parse a single line of input"""
        line = line.strip()
        if not line:
            return None
            
        if self.is_valid_transition(line):
            states = re.findall(self.state_pattern, line)
            return f"Valid transition from {states[0]} to {states[1]}"
        else:
            return f"Invalid input: {line}"

def main():
    parser = Parser()
    # Read from stdin
    for line in sys.stdin:
        result = parser.parse_line(line)
        if result:
            print(result)

if __name__ == "__main__":
    main()
