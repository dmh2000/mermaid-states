# Mermaid State Diagram To Code

This project has two components:
- A state machine library, for go, c++ and c languages.
- A mermaid state diagram command line tool that converts a mermaid state diagram to code that implements scaffolds for state handlers that are compatible with the state machine library. The output code can be generated as go, c++ or c code.

The output of the mermaid state diagram tool can be used to generate the state machine library code for the particular language. 

## Mermaid State Diagram Tool

The syntax for a mermaid state diagram is defined in the [mermaid](https://mermaid.js.org/syntax/stateDiagram.html) documentation. The tool will accept a mermaid file that contains state diagram components. It ignores any other mermaid syntax. 

## State Machine Library

The state machine library is a generic state machine implementation that can be used for any language. It is implemented in go, c++ and c. In this library, a state machine is composed of 4 generic components:
- Model: a generic model that is passed to the state machine. It can be any type. The 'Model' type is used to store persistent data that the state machines operates on.

- Input: a generic input that is passed to the state machine. It can be any type. The 'Input' type represents extern inputs that drive the state machine..
- State: a generic state that is used by the state machine. A state object looks like this:

```go
// StateKey represents a unique key for a state in the state machine.
type StateKey string

// ActionFunc is called when a state is executed.
type ActionFunc[Model any, Input any] func(model *Model, input Input) (key StateKey, err error)

// State represents a state in the state machine with a key, name, and action.
type State[Model any, Input any] struct {
	Key    StateKey
	action ActionFunc[Model, Input]
}
```

```c++
// State represents a state in the state machine with a key, name, and action.
struct State {
    StateKey key;
    ActionFunc action;
};
```
- Key: a string that identifies a state. In go it is type aliased as StateKey
- ActionFunc: a function that is called when a state is executed. It takes a model and an input and returns a state key and an error.



## Code structure

The code is divided into two subdirectories:

- [go/parser](go/parser): contains the code for parsing mermaid state diagram syntax. The code is in the [parser.go](go/parser/parser.go) file.

- [go/graph](go/graph): contains the code for creating a graph of states and transitions. The code is in the [graph.go](go/graph/graph.go) file.