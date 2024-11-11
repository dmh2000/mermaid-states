# regexes for state, description, and transition

- State
  - ^(?:[A-Za-z\_][A-Za-z0-9_]\*|\[\*\])$
  - State names must start with letter or underscore, followed by letter, digit, or underscore or [\*]
- Transition
  - ^([A-Za-z\_][A-Za-z0-9_]_|\[\*\])\s_-->\s*([A-Za-z\_][A-Za-z0-9_]*|\[\*\])(?:\s\*\:(.+))?$
  - Transitions must start with state name, followed by --> and another state name and optional description
    // descriptions must start with colon, followed by one or
    more characters
- Description
  - ^.+$
  - any printable characters
