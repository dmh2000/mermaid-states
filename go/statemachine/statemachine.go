// Package statemachine provides a simple state machine implementation
package statemachine

import (
	"fmt"
)

// StateKey represents a unique key for a state in the state machine.
type StateKey int32

// String returns the string representation of the state key.
func (k StateKey) String() string {
	return fmt.Sprintf("%d", k)
}

// State represents a state in the state machine with a key, name, and action.
type State[Model any, Input any] struct {
	key    StateKey
	name   string
	action func(*Model, Input) (key StateKey, err error)
}

// String returns the string representation of the state.
func (s *State[Model, Input]) String() string {
	return fmt.Sprintf("key: %d , name: %s", s.key, s.name)
}

// GetKey returns the key of the state.
func (s *State[Model, Input]) GetKey() StateKey {
	return s.key
}

// GetName returns the name of the state.
func (s *State[Model, Input]) GetName() string {
	return s.name
}

// Execute performs the state's action and returns the key of the next state.
func (s *State[Model, Input]) Execute(model *Model, input Input) (key StateKey, err error) {
	return s.action(model, input)
}

// NewState creates a new state with the given key, name, and action.
func NewState[Model any, Input any](
	key StateKey,
	name string,
	act func(*Model, Input) (key StateKey, err error),
) *State[Model, Input] {
	if act == nil {
		panic("action function cannot be nil")
	}
	return &State[Model, Input]{
		key:    key,
		name:   name,
		action: act,
	}
}

// ============================================================================

// StateMachine represents a state machine with a current state and a collection of states.
type StateMachine[Model any, Input any] struct {
	currentState *State[Model, Input]
	states       map[StateKey]*State[Model, Input]
	name         string
}

// NewStateMachine creates a new state machine with the given model and name.
func NewStateMachine[Model any, Input any](model *Model, name string) *StateMachine[Model, Input] {
	return &StateMachine[Model, Input]{
		currentState: nil,
		states:       make(map[StateKey]*State[Model, Input]),
		name:         name,
	}
}

// String returns the string representation of the state machine.
func (sm *StateMachine[Model, Input]) String() string {
	return fmt.Sprintf("name: %s , states: %d", sm.name, len(sm.states))
}

// GetCurrentState returns the current state of the state machine.
func (sm *StateMachine[Model, Input]) GetCurrentState() *State[Model, Input] {
	return sm.currentState
}

// GetStates returns a map of all available states.
func (sm *StateMachine[Model, Input]) GetStates() map[StateKey]*State[Model, Input] {
	states := make(map[StateKey]*State[Model, Input])
	for k, v := range sm.states {
		states[k] = v
	}
	return states
}

// AddState adds a new state to the state machine. If it's the first state, it sets it as the initial state.
func (sm *StateMachine[Model, Input]) AddState(state *State[Model, Input]) error {
	// check if state already exists
	if _, exists := sm.states[state.key]; exists {
		return fmt.Errorf("state %d already exists", state.key)
	}
	// ok, add it to the map
	sm.states[state.key] = state

	// if there is no current state, set it as the initial state
	if sm.currentState == nil {
		if err := sm.SetInitialState(state.key); err != nil {
			return err
		}
	}
	return nil
}

// SetInitialState sets the initial state of the state machine using the given key.
func (sm *StateMachine[Model, Input]) SetInitialState(key StateKey) error {

	state, exists := sm.states[key]
	if !exists {
		return fmt.Errorf("state %d does not exist", key)
	}

	sm.currentState = state

	return nil
}

// Execute performs the current state's action and transitions to the next state based on the returned key.
func (sm *StateMachine[Model, Input]) Execute(model *Model, input Input) (key StateKey, err error) {
	if sm.currentState == nil {
		return 0, fmt.Errorf("no current state set")
	}
	
	key, err = sm.currentState.Execute(model, input)
	if err != nil {
		return
	}

	// same state, no change
	if key == sm.currentState.GetKey() {
		return key, nil
	}

	// new state
	newState, exists := sm.states[key]
	if !exists {
		return 0, fmt.Errorf("state %d does not exist", key)
	}

	// set next state
	sm.currentState = newState

	return newState.GetKey(), nil
}
