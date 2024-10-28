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
    
    def parse_composite_state(self, lines, indent=0):
        """Parse a composite state block"""
        results = []
        i = 0
        while i < len(lines):
            line = lines[i].strip()
            if not line:
                i += 1
                continue
                
            if line.startswith("state "):
                # Extract state name and verify it's valid
                state_name = line.split()[1]
                if not self.is_valid_state(state_name):
                    results.append("  " * indent + f"Invalid state name: {state_name}")
                    i += 1
                    continue
                    
                # Skip the opening brace
                i += 1
                if i >= len(lines) or lines[i].strip() != "{":
                    results.append("  " * indent + "Missing opening brace after state declaration")
                    continue
                results.append("  " * indent + f"Composite state {state_name}:")
                nested_lines = []
                brace_count = 1
                i += 1  # Move past the opening brace
                while i < len(lines) and brace_count > 0:
                    if lines[i].strip() == "{":
                        brace_count += 1
                    elif lines[i].strip() == "}":
                        brace_count -= 1
                    if brace_count > 0:
                        nested_lines.append(lines[i])
                    i += 1
                if not nested_lines:
                    results.append("  " * indent + "Error: Composite state must contain at least one transition")
                else:
                    # Parse nested lines and check if there's at least one valid transition
                    nested_results = self.parse_composite_state(nested_lines, indent + 1)
                    has_transition = any("Valid transition from" in result for result in nested_results)
                    
                    if not has_transition:
                        results.append("  " * indent + "Error: Composite state must contain at least one transition")
                    else:
                        results.append("  " * indent + "Composite state contains:")
                        results.extend(nested_results)
            elif self.is_valid_transition(line):
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
