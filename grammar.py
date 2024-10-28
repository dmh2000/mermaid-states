import sys
import re

class Parser:
    def __init__(self):
        # Define regex patterns based on BNF grammar
        self.state_pattern = r'[A-Za-z_][A-Za-z0-9_]*'
        self.description_pattern = r':[a-zA-Z0-9_\-:.,?!@=~ ]+'
        self.transition_pattern = rf'{self.state_pattern}\s*-->\s*{self.state_pattern}(?:\s*{self.description_pattern})?$'
        
    def is_valid_state(self, state):
        """Check if a string matches the STATE rule"""
        return bool(re.match(f'^{self.state_pattern}$', state))
    
    def is_valid_transition(self, line):
        """Check if a string matches the TRANSITION rule"""
        return bool(re.match(f'^{self.transition_pattern}$', line))
    
    def parse_graph(self, lines, indent=0):
        """Parse a graph consisting of one or more transitions"""
        results = []
        i = 0
        while i < len(lines):
            line = lines[i].strip()
            if not line:
                i += 1
                continue
                
            if self.is_valid_transition(line):
                # Extract states and description
                match = re.match(rf'^({self.state_pattern})\s*-->\s*({self.state_pattern})(?:\s*({self.description_pattern}))?$', line)
                if match:
                    from_state, to_state, description = match.groups()
                    desc_text = f" with description '{description[1:].strip()}'" if description else ""
                    results.append("  " * indent + f"Valid transition from {from_state} to {to_state}{desc_text}")
            else:
                results.append("  " * indent + f"Invalid input: {line}")
            i += 1
        return results

    def parse_input(self):
        """Parse all input as a graph of transitions"""
        lines = []
        for line in sys.stdin:
            lines.append(line.rstrip())
        
        # Verify there is at least one transition
        valid_lines = [line for line in lines if line.strip() and self.is_valid_transition(line.strip())]
        if not valid_lines:
            return ["Error: Graph must contain at least one transition"]
            
        return self.parse_graph(lines)

def main():
    parser = Parser()
    results = parser.parse_input()
    for result in results:
        print(result)

if __name__ == "__main__":
    main()
