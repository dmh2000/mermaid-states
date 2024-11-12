// Package statemachine provides a simple state machine implementation
package statemachine

import (
	"fmt"
)

/**
Model type : a pointer to the common data that is shared between states
Input type : the input data that is passed to the state
*/

type StateKey int32

type State[Model any, Input any] struct {
	key    StateKey
	name   string
	action func(*Model, Input) (key int32, err error)
}

func (s *State[Model, Input]) String() string {
	return fmt.Sprintf("State %d: %s", s.key, s.name)
}

func (s *State[Model, Input]) GetKey() StateKey {
	return s.key
}

func (s *State[Model, Input]) GetName() string {
	return s.name
}

// Execute :  the state action and return a key to the next state (could be the same or different)
func (s *State[Model, Input]) Execute(model *Model, input Input) (key int32, err error) {
	return s.action(model, input)
}

func NewState[Model any, Input any](
	key StateKey,
	name string,
	act func(*Model, Input) (key int32, err error),
) *State[Model, Input] {
	return &State[Model, Input]{
		key:    key,
		name:   name,
		action: act,
	}
}

// ============================================================================

type StateMachine[Model any, Input any] struct {
	currentState *State[*Model, Input]
	states       map[StateKey]*State[Model, Input]
	name         string
}

func NewStateMachine[Model any, Input any](model *Model, name string) *StateMachine[Model, Input] {
	return &StateMachine[Model, Input]{
		currentState: nil,
		states:       make(map[StateKey]*State[Model, Input]),
		name:         name,
	}
}

func (sm *StateMachine[Model, Input]) String() string {
	return fmt.Sprintf("StateMachine %s , %d", sm.name, len(sm.states))
}

func (sm *StateMachine[Model, Input]) AddState(state *State[Model, Input]) error {
	// check if state already exists
	if _, exists := sm.states[state.key]; exists {
		return fmt.Errorf("state %d already exists", state.key)
	}
	// add it to the map
	sm.states[state.key] = state
	return nil
}

// ============================================================================

// func AddState[data Data](sm *StateMachine[data], state *State[data]) error {

// 	// check if state already exists
// 	if _, exists := sm.states[state.Key]; exists {
// 		return fmt.Errorf("state %d already exists", state.Key)
// 	}
// 	// add it to the map
// 	sm.states[state.Key] = state

// 	// if its the first state, set it as the initial state
// 	if sm.current_state == nil {
// 		if err := SetInitialState(sm, state.Key); err != nil {
// 			return err
// 		}
// 	}

// 	// no errors
// 	return nil
// }

// func SetInitialState[data Data](sm *StateMachine[data], key int32) error {
// 	state, exists := sm.states[key]
// 	if !exists {
// 		return fmt.Errorf("state %d does not exist", key)
// 	}
// 	sm.current_state = state
// 	return nil
// }

// func executeState[data Data](sm *StateMachine[data], input Input) error {
// 	if sm.current_state == nil {
// 		return fmt.Errorf("no current state set")
// 	}

// 	state, exists := sm.states[sm.current_state.Key]
// 	if !exists {
// 		return fmt.Errorf("no current state set")
// 	}

// 	next_state, err := state.Action(state, input)
// 	if err != nil {
// 		return err
// 	}

// 	sm.current_state = next_state
// 	return nil
// }

// func Run[data Data](sm *StateMachine[data], input Input) error {
// 	return executeState(sm, input)
// }
