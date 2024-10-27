import sys
import re

class Parser:
    def __init__(self):
        # Define regex patterns based on BNF grammar
        self.state_pattern = r'[A-Za-z_][A-Za-z0-9_]*'
        self.description_pattern = r':\s*[A-Za-z0-9]+'
        self.transition_pattern = rf'{self.state_pattern}\s*-->\s*{self.state_pattern}(?:\s*{self.description_pattern})?'
        
    def is_valid_state(self, state):
        """Check if a string matches the STATE rule"""
        return bool(re.match(f'^{self.state_pattern}$', state))
    
    def is_valid_transition(self, line):
        """Check if a string matches the TRANSITION rule"""
        return bool(re.match(f'^{self.transition_pattern}$', line))
    
    def parse_composite_state(self, lines, indent=0):
        """Parse a composite state block"""
        results = []
        i = 0
        while i < len(lines):
            line = lines[i].strip()
            if not line:
                i += 1
                continue
                
            if line == "{":
                nested_lines = []
                brace_count = 1
                i += 1
                while i < len(lines) and brace_count > 0:
                    if lines[i].strip() == "{":
                        brace_count += 1
                    elif lines[i].strip() == "}":
                        brace_count -= 1
                    if brace_count > 0:
                        nested_lines.append(lines[i])
                    i += 1
                if nested_lines:
                    results.append("  " * indent + "Composite state contains:")
                    results.extend(self.parse_composite_state(nested_lines, indent + 1))
            elif self.is_valid_transition(line):
                states = re.findall(self.state_pattern, line)
                description = re.findall(self.description_pattern, line)
                desc_text = f" with description '{description[0][1:].strip()}'" if description else ""
                results.append("  " * indent + f"Valid transition from {states[0]} to {states[1]}{desc_text}")
            else:
                results.append("  " * indent + f"Invalid input: {line}")
            i += 1
        return results

    def parse_input(self):
        """Parse all input and handle composite states"""
        lines = []
        for line in sys.stdin:
            lines.append(line.rstrip())
        return self.parse_composite_state(lines)

def main():
    parser = Parser()
    results = parser.parse_input()
    for result in results:
        print(result)

if __name__ == "__main__":
    main()
