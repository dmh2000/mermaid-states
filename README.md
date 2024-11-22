# Mermaid State Diagram To Code

This project has two components:

- A state machine library, for go, c++ and c languages.
- A mermaid state diagram command line tool that converts a mermaid state diagram to code that implements scaffolds for state handlers that are compatible with the state machine library. The output code can be generated as go, c++ or c code.

The output of the mermaid state diagram tool can be used to generate the state machine library code for the particular language.

## Mermaid State Diagram Tool

The syntax for a mermaid state diagram is defined in the [mermaid](https://mermaid.js.org/syntax/stateDiagram.html) documentation. The tool will accept a mermaid file that contains state diagram components. It ignores any other mermaid syntax.

The output of the state diagram tool is a code file that creates the states , adds them to the state machine and executes the state machine. The state action functions are scaffolded with comments. Its up to the developer to flesh out the individual states actions.

## State Machine Library

The state machine library is a generic state machine library implementation that can be used for supported languages. It is implemented in go, c++ and c. For this library, a state machine is composed of 4 generic components:

- **Model**: a generic model that is passed to the state machine. It can be any type. The 'Model' type is used to store persistent data that the state machines operates on.
- **Input**: a generic input that is passed to the state machine. It can be any type. The 'Input' type represents extern inputs that drive the state machine..
- **State**: a generic state that is used by the state machine. A state object looks like this:
- **Action**: a generic function that is called when a state is executed. It receives a pointer to the current state, a pointer to the 'global' model, and the input. The Model and Input types are generic.

### GO

- Using the StateMachine Library

```go
// go/pkg/statemachine
// StateKey represents a unique key for a state in the state machine.
type StateKey string

// ActionFunc is called when a state is executed.
type ActionFunc[Model any, Input any] func(currentState *State[Model, Input], model *Model, input Input) (key StateKey, err error)

// State represents a state in the state machine with a key, name, and action.
type State[Model any, Input any] struct {
	Key    StateKey                     // a string used to identify a state in the state map
	Action ActionFunc[Model, Input]     // a function that is called when a state is executed
    Data   *interface{}                 // a pointer to arbitrary data that is associated with the specific state
}

// StateMachine represents a state machine with a current state and a collection of states.
type StateMachine[Model any, Input any] struct {
	currentState *State[Model, Input]
	states       map[StateKey]*State[Model, Input]
	name         string
}

```

#### Definie States With the NewState Function

- specify a key (string)
- specify an action function (ActionFunc)
- specify optional data (interface{})

```go
// inline action functions
state1 := NewState(
    "state1",
	func (currentState *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
		model.value = input
		// ....
		return s2, nil
	},
    nil,
)

// separate action function
func state1Action(currentState *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
	model.value = input
	// ....
	return s2, nil
}
// ...
state1 := NewState(
    "state1",
	state1Action,
    nil,
)
```

#### Instantiate a state machine, add states and execute it

```go
// go/pkg/statemachine

// create the state machine
sm := NewStateMachine[testModel, int](
    &testModel{0},
    "test",
)

state1 := NewState(
    "state1",
	func (s *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
		model.value = input
		// ....
		return s2, nil
	},
    nil,
)

// add the state
if err := sm.AddState(state1); err != nil {
	t.Errorf("unexpected error: %s", err)
}

// more states...

// optionally set the initial state (otherwise it will be the first state added)
sm.SetInitialState(s2)

// create the model
model := &testModel{0}

// execute the state machine until its done
for {
	input := rand.Intn(100) // whatever the input is

	key, err := sm.Execute(&model, input)
	if err != nil {
		t.Fatalf("unexpected error during execution: %s", err)
	}

	if key == "END" {
		break
	}
}

```

### C++

### C

## Code structure

The code is divided into two subdirectories:

- [go/parser](go/parser): contains the code for parsing mermaid state diagram syntax. The code is in the [parser.go](go/parser/parser.go) file.

- [go/graph](go/graph): contains the code for creating a graph of states and transitions. The code is in the [graph.go](go/graph/graph.go) file.
